/*
 * Copyright 2020, Offchain Labs, Inc.
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

#ifndef cnodestore_h
#define cnodestore_h

#include "ctypes.h"

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

void deleteConfirmedNodeStore(CConfirmedNodeStore* m);
int putNode(CConfirmedNodeStore* node_store,
            uint64_t height,
            const void* hash,
            const void* data,
            int data_length);
ByteSliceResult getNode(CConfirmedNodeStore* node_store,
                        uint64_t height,
                        const void* hash);
Uint64Result getNodeHeight(CConfirmedNodeStore* node_store, const void* hash);
HashResult getNodeHash(CConfirmedNodeStore* node_store, uint64_t height);
int isNodeStoreEmpty(CConfirmedNodeStore* node_store);
uint64_t maxNodeHeight(CConfirmedNodeStore* node_store);

#ifdef __cplusplus
}
#endif

#endif /* cnodestore_h */
