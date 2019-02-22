/*
 *    Copyright 2019 Insolar Technologies
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package storage

// func TestObjectStorageMEM_SetRecord(t *testing.T) {
// 	t.Parallel()
//
// 	// Arrange
// 	ctx := inslogger.TestContext(t)
//
// 	objectStorage := &objectStorageMEM{
// 		recordStorage: objectsPerJet{},
// 	}
//
// 	jetID := testutils.RandomJet()
//
// 	rec := &record.RequestRecord{}
// 	_, err := objectStorage.SetRecord(ctx, jetID, core.GenesisPulse.PulseNumber, rec)
//
// 	assert.NoError(t, err)
// 	assert.Equal(t, 1, len(objectStorage.recordStorage))
// 	assert.Equal(t, 1, objectStorage.recordStorage.pulseSize)
// 	assert.Equal(t, 1, objectStorage.recordStorage.memSize)
//
// 	assert.True(t, false)
// }