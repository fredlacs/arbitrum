/*
 * Copyright 2019, Offchain Labs, Inc.
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

pragma solidity ^0.5.3;

import "./Value.sol";


library Protocol {
    using Value for Value.Data;

    function generateMessageStubHash(
        Value.Data memory _dataHashValue,
        bytes21 _tokenType,
        uint256 _value,
        address _destination
    )
        internal
        pure
        returns (bytes32)
    {
        Value.Data[] memory values = new Value.Data[](4);
        values[0] = _dataHashValue;
        values[1] = Value.newInt(uint256(_destination));
        values[2] = Value.newInt(_value);
        values[3] = Value.newInt(uint256(bytes32(_tokenType)));
        return Value.newTuple(values).hash();
    }

    function generatePreconditionHash(
        bytes32 _beforeHash,
        uint128[4] memory _timeBounds,
        bytes32 _beforeInboxHash
    )
        internal
        pure
        returns (bytes32)
    {
        return keccak256(
            abi.encodePacked(
                _beforeHash,
                _timeBounds[0],
                _timeBounds[1],
                _timeBounds[2],
                _timeBounds[3],
                _beforeInboxHash
            )
        );
    }

    function generateAssertionHash(
        bytes32 _afterHash,
        bool    _didInboxInsn,
        uint64  _numGas,
        bytes32 _firstMessageHash,
        bytes32 _lastMessageHash,
        bytes32 _firstLogHash,
        bytes32 _lastLogHash
    )
        internal
        pure
        returns (bytes32)
    {
        return keccak256(
            abi.encodePacked(
                _afterHash,
                _didInboxInsn,
                _numGas,
                _firstMessageHash,
                _lastMessageHash,
                _firstLogHash,
                _lastLogHash
            )
        );
    }
}
