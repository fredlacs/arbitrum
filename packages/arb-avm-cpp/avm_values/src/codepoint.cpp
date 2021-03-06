/*
 * Copyright 2019-2020, Offchain Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

#include <avm_values/codepoint.hpp>
#include <avm_values/tuple.hpp>
#include <avm_values/util.hpp>
#include <bigint_utils.hpp>

Operation::Operation(OpCode opcode_, value immediate_)
    : opcode(opcode_), immediate(std::make_unique<value>(immediate_)) {}

Operation::Operation(const Operation& op) {
    opcode = op.opcode;
    if (op.immediate) {
        immediate = std::make_unique<value>(*op.immediate);
    }
}
Operation::Operation(Operation&&) = default;
Operation& Operation::operator=(const Operation& cp) {
    opcode = cp.opcode;
    if (cp.immediate) {
        immediate = std::make_unique<value>(*cp.immediate);
    } else {
        immediate.reset();
    }
    return *this;
}

Operation& Operation::operator=(Operation&&) = default;
Operation::~Operation() = default;

void Operation::marshal(std::vector<unsigned char>& buf,
                        const Code& code) const {
    if (immediate) {
        buf.push_back(1);
        buf.push_back((uint8_t)opcode);
        marshal_value(*immediate, buf, code);
    } else {
        buf.push_back(0);
        buf.push_back((uint8_t)opcode);
    }
}

void Operation::marshalForProof(std::vector<unsigned char>& buf,
                                bool includeVal,
                                const Code& code) const {
    if (immediate) {
        buf.push_back(1);
        buf.push_back((uint8_t)opcode);
        if (includeVal) {
            ::marshalForProof(*immediate, buf, code);
        } else {
            marshalStub(*immediate, buf, code);
        }
    } else {
        buf.push_back(0);
        buf.push_back((uint8_t)opcode);
    }
}

uint64_t pc_default = -1;

bool operator==(const CodePoint& val1, const CodePoint& val2) {
    if (val1.pc != val2.pc)
        return false;
    else
        return true;
}

void CodePoint::marshal(std::vector<unsigned char>& buf,
                        const Code& code) const {
    buf.push_back(CODEPT);
    uint64_t bepc = boost::endian::native_to_big(pc);
    std::copy(
        static_cast<const char*>(static_cast<const void*>(&bepc)),
        static_cast<const char*>(static_cast<const void*>(&bepc)) + sizeof bepc,
        std::back_inserter(buf));
    op.marshal(buf, code);
    std::array<unsigned char, 32> val;
    to_big_endian(nextHash, val.begin());
    buf.insert(buf.end(), val.begin(), val.end());
}

uint256_t hash(const CodePoint& cp) {
    std::array<uint64_t, 4> nextHashInts;
    to_big_endian(cp.nextHash, nextHashInts.begin());
    if (cp.op.immediate) {
        std::array<unsigned char, 66> valData;
        valData[0] = CODEPT;
        valData[1] = static_cast<unsigned char>(cp.op.opcode);
        auto immHash = hash_value(*cp.op.immediate);
        std::array<uint64_t, 4> valHashInts;
        to_big_endian(immHash, valHashInts.begin());
        std::copy(reinterpret_cast<unsigned char*>(valHashInts.data()),
                  reinterpret_cast<unsigned char*>(valHashInts.data()) + 32,
                  valData.begin() + 2);
        std::copy(reinterpret_cast<unsigned char*>(nextHashInts.data()),
                  reinterpret_cast<unsigned char*>(nextHashInts.data()) + 32,
                  valData.end() - 32);
        std::array<unsigned char, 32> hashData;
        evm::Keccak_256(valData.data(), valData.size(), hashData.data());
        return from_big_endian(hashData.begin(), hashData.end());
    } else {
        std::array<unsigned char, 34> valData;
        valData[0] = CODEPT;
        valData[1] = static_cast<unsigned char>(cp.op.opcode);
        std::copy(reinterpret_cast<unsigned char*>(nextHashInts.data()),
                  reinterpret_cast<unsigned char*>(nextHashInts.data()) + 32,
                  valData.end() - 32);
        std::array<unsigned char, 32> hashData;
        evm::Keccak_256(valData.data(), valData.size(), hashData.data());
        return from_big_endian(hashData.begin(), hashData.end());
    }
}

std::ostream& operator<<(std::ostream& os, const Operation& val) {
    if (val.immediate) {
        os << "Immediate(" << InstructionNames.at(val.opcode) << ", "
           << *val.immediate << ")";
    } else {
        os << "Basic(" << InstructionNames.at(val.opcode) << ")";
    }
    return os;
}

std::ostream& operator<<(std::ostream& os, const CodePoint& val) {
    os << "CodePoint(" << val.pc << ", " << val.op << ", "
       << to_hex_str(val.nextHash) << ")";
    return os;
}

const CodePoint& getErrCodePoint() {
    CodePoint static errcp(pc_default, Operation(static_cast<OpCode>(0)), 0);
    return errcp;
}

std::vector<CodePoint> opsToCodePoints(const std::vector<Operation>& ops) {
    std::vector<CodePoint> cps;
    cps.reserve(ops.size());
    uint64_t pc = 0;
    for (auto& op : ops) {
        cps.emplace_back(pc, std::move(op), 0);
        pc++;
    }
    for (uint64_t i = 0; i < cps.size() - 1; i++) {
        cps[cps.size() - 2 - i].nextHash = hash(cps[cps.size() - 1 - i]);
    }
    return cps;
}

Code::Code(std::vector<CodePoint> code_) : code(std::move(code_)) {}

std::ostream& operator<<(std::ostream& os, const Code& code) {
    for (const auto& cp : code.code) {
        os << cp << "\n";
    }
    return os;
}
