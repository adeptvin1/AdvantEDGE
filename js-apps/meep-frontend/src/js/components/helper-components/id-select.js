/*
 * Copyright (c) 2019  InterDigital Communications, Inc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import React from 'react';
import { Select } from '@rmwc/select';
import { GridCell } from '@rmwc/grid';

const IDSelect = props => {
  return (
    <GridCell span={props.span}>
      <Select
        style={{ width: '100%' }}
        label={props.label}
        outlined={props.isDialog?false:true}
        options={props.options}
        onChange={props.onChange}
        disabled={props.disabled}
        value={props.value}
        data-cy={props.cydata}
      />
    </GridCell>
  );
};

export default IDSelect;
