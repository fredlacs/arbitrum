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

#ifndef checkpoint_value_hpp
#define checkpoint_value_hpp

#include <avm_values/value.hpp>

struct DeleteResults;
struct SaveResults;
class Transaction;

template <typename T>
struct DbResult;

DbResult<value> getValue(const Transaction& transaction,
                         uint256_t value_hash,
                         TuplePool* pool);
SaveResults saveValue(Transaction& transaction, const value& val);
DeleteResults deleteValue(Transaction& transaction, uint256_t value_hash);

#endif /* value_hpp */
