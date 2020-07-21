/* -------------------------------------------------------------------------------------/
                                                                                    /
/               COPYRIGHT (c) 2019 SAMSUNG ELECTRONICS CO., LTD.                      /
/                          ALL RIGHTS RESERVED                                        /
/                                                                                     /
/   Permission is hereby granted to licensees of Samsung Electronics Co., Ltd.        /
/   products to use or abstract this computer program for the sole purpose of         /
/   implementing a product based on Samsung Electronics Co., Ltd. products.           /
/   No other rights to reproduce, use, or disseminate this computer program,          /
/   whether in part or in whole, are granted.                                         / 
/                                                                                     /
/   Samsung Electronics Co., Ltd. makes no representation or warranties with          /
/   respect to the performance of this computer program, and specifically disclaims   /
/   any responsibility for any damages, special or consequential, connected           /
/   with the use of this program.                                                     /
/                                                                                     /
/-------------------------------------------------------------------------------------/


DESCRIPTION: Alert Management Container Test File
@NAME : index.test.js
@AUTHORS: Jay Hitesh Sanghavi 
@Version : 1.0 *
@REVISION HISTORY
[03/06/2019] [Jay] : Prototyping..........////////////////////
*/
import React from 'react';
import { Router } from 'react-router-dom';
import { I18nextProvider } from "react-i18next";
import { combineReducers, createStore, applyMiddleware, compose } from 'redux'
import createSagaMiddleware from 'redux-saga'
import MockAdapter from 'axios-mock-adapter'
import { act } from 'react-dom/test-utils'
import { Provider } from 'react-redux'
import axios from 'axios'
import { createMemoryHistory } from 'history';
import { render, fireEvent, cleanup, waitForElement } from '@testing-library/react';
import AlertManagement from './index';
import headerReducer from "../../../store/reducers/headerReducer"
import headerLanguageReducer from "../../../store/reducers/headerLanguageReducer"
import alertManagementReducer from "../../../store/reducers/alertManagementReducer"
import configurationsettingReducer from "../../../store/reducers/configurationsettingReducer"
import rootSaga from "../../../sagas/indexSaga"
import i18n from "../../../i18n";

jest.unmock('axios');

describe("Alert Management", () => {
  let select_input;
  let select_save;
  let select_delete;
  let select_switch_tag;
  let select_input_child;
  let wrapper;
  let store;
  let history;
  let alerts = [{ "_id": { "$oid": "5d5e67b7f082d4c68c98f58b" }, "alertName": "NewAlert", "alertCluster": "cpu", "alertSubCluster": "device", "alertType": "cpu-total", "alertCondition": null, "alertField": "usage_system", "description": "last", "alertRange": "109", "active": true },]
  let alertClusterList = [{ alertFields: ['usage_idle', 'usage_system', 'NA'] }];
  const radioindex = 0;
  const alertClusterName = "Random";
  const selectedAlertSubCluster = 'Random';
  const alertType = 'Random';

  beforeEach(() => {
    const sagaMiddleware = createSagaMiddleware();
    const composeEnhancers = window.REDUX_DEVTOOLS_EXTENSION_COMPOSE || compose;

    const rootReducers = combineReducers({
      alertManagementReducer,
      headerReducer,
      headerLanguageReducer,
      configurationsettingReducer
    });
    store = createStore(rootReducers, composeEnhancers(applyMiddleware(sagaMiddleware)))

    sagaMiddleware.run(rootSaga);
    const route = '/';
    history = createMemoryHistory({ initialEntries: [route] })
    wrapper = render(
      <Router history={history}>
        <I18nextProvider i18n={i18n}>
          <Provider store={store}>
            <AlertManagement alerts={alerts} />
          </Provider>
        </I18nextProvider>
      </Router>
    );
  });
  afterEach(cleanup);

  it('renders alert management saga', async () => {
    const mock = new MockAdapter(axios);
    let response = mock.onGet('/api/v1.0/get_alerts/').reply(200, null);
    response = mock.onGet('/api/v1.0/get_alert_types/').reply(200, null);
  });

  it('renders alert management saga', async () => {
    const mock = new MockAdapter(axios);
    let response = mock.onGet('/api/v1.0/get_alerts/').reply(401, null);
    response = mock.onGet('/api/v1.0/get_alert_types/').reply(401, null);
  });

  it('renders alert management saga', async () => {
    const mock = new MockAdapter(axios);
    let data = [{ "_id": { "$oid": "5d5e67b7f082d4c68c98f58b" }, "alertName": "NewAlert", "alertCluster": "cpu", "alertSubCluster": "device", "alertType": "cpu-total", "alertCondition": "Greater Than", "alertField": "usage_system", "description": "last", "alertRange": "109", "active": true }, { "_id": { "$oid": "5d660a4bd4bf85b0c852ed03" }, "alertName": "ThisisAwesome", "alertCluster": "cpu", "alertSubCluster": "cpu ", "alertType": "cpu-total", "alertCondition": "Equal To", "alertField": "usage_user", "description": "Hmmm", "alertRange": "93", "active": true }, { "_id": { "$oid": "5d666e0bd4bf85b0c852ed05" }, "alertName": "Alert1", "alertCluster": "cpu", "alertSubCluster": "cpu ", "alertType": "cpu-total", "alertCondition": "Equal To Or Greater", "alertField": "usage_user", "description": "Yayyy", "alertRange": "34", "active": true }];
    let response = mock.onGet('/api/v1.0/get_alerts/').reply(200, data);
    data = {
      "alert_types": [
        {
          "_id": 1,
          "alertFields": [
            "lat_data_0_id",
            "lat_data_0_tid_arr_0_aid_arr_0_aid",
            "lat_data_0_tid_arr_0_aid_arr_0_cnt_miss_low",
            "lat_data_0_tid_arr_0_aid_arr_0_cnt_miss_up",
            "lat_data_0_tid_arr_0_aid_arr_0_cnt_sample",
            "lat_data_0_tid_arr_0_aid_arr_0_low_qt",
            "lat_data_0_tid_arr_0_aid_arr_0_max",
            "lat_data_0_tid_arr_0_aid_arr_0_mean",
            "lat_data_0_tid_arr_0_aid_arr_0_min",
            "lat_data_0_tid_arr_0_aid_arr_0_phase",
            "lat_data_0_tid_arr_0_aid_arr_0_qos_2nine",
            "lat_data_0_tid_arr_0_aid_arr_0_qos_3nine",
            "lat_data_0_tid_arr_0_aid_arr_0_qos_4nine",
            "lat_data_0_tid_arr_0_aid_arr_0_std",
            "lat_data_0_tid_arr_0_aid_arr_0_up_qt",
            "lat_data_0_tid_arr_0_tid",
            "lat_data_0_tid_arr_1_aid_arr_0_aid",
            "lat_data_0_tid_arr_1_aid_arr_0_cnt_miss_low",
            "lat_data_0_tid_arr_1_aid_arr_0_cnt_miss_up",
            "lat_data_0_tid_arr_1_aid_arr_0_cnt_sample",
            "lat_data_0_tid_arr_1_aid_arr_0_low_qt",
            "lat_data_0_tid_arr_1_aid_arr_0_max",
            "lat_data_0_tid_arr_1_aid_arr_0_mean",
            "lat_data_0_tid_arr_1_aid_arr_0_min",
            "lat_data_0_tid_arr_1_aid_arr_0_phase",
            "lat_data_0_tid_arr_1_aid_arr_0_qos_2nine",
            "lat_data_0_tid_arr_1_aid_arr_0_qos_3nine",
            "lat_data_0_tid_arr_1_aid_arr_0_qos_4nine",
            "lat_data_0_tid_arr_1_aid_arr_0_std",
            "lat_data_0_tid_arr_1_aid_arr_0_up_qt",
            "lat_data_0_tid_arr_1_tid",
            "perf_data_0_id",
            "perf_data_0_tid_arr_0_tid",
            "perf_data_0_tid_arr_1_aid_arr_0_aid",
            "perf_data_0_tid_arr_1_aid_arr_0_bw_read",
            "perf_data_0_tid_arr_1_aid_arr_0_bw_total",
            "perf_data_0_tid_arr_1_aid_arr_0_bw_write",
            "perf_data_0_tid_arr_1_aid_arr_0_cnt_128kb",
            "perf_data_0_tid_arr_1_aid_arr_0_cnt_16kb",
            "perf_data_0_tid_arr_1_aid_arr_0_cnt_1kb",
            "perf_data_0_tid_arr_1_aid_arr_0_cnt_256kb",
            "perf_data_0_tid_arr_1_aid_arr_0_cnt_2kb",
            "perf_data_0_tid_arr_1_aid_arr_0_cnt_32kb",
            "perf_data_0_tid_arr_1_aid_arr_0_cnt_4kb",
            "perf_data_0_tid_arr_1_aid_arr_0_cnt_512b",
            "perf_data_0_tid_arr_1_aid_arr_0_cnt_64kb",
            "perf_data_0_tid_arr_1_aid_arr_0_cnt_8kb",
            "perf_data_0_tid_arr_1_aid_arr_0_cnt_big",
            "perf_data_0_tid_arr_1_aid_arr_0_cnt_no_align",
            "perf_data_0_tid_arr_1_aid_arr_0_iops_read",
            "perf_data_0_tid_arr_1_aid_arr_0_iops_total",
            "perf_data_0_tid_arr_1_aid_arr_0_iops_write",
            "perf_data_0_tid_arr_1_tid",
            "streaming_interval"
          ],
          "alertSubCluster": [
            {
              "_id": 1,
              "alertTypes": [
                {
                  "type": "2030045418"
                }
              ],
              "name": "host"
            },
            {
              "_id": 2,
              "alertTypes": [
                {
                  "type": "/home/dinesh/ibof_devel_1206/ibofos/air_result.json"
                }
              ],
              "name": "path"
            }
          ],
          "name": "air"
        },
        {
          "_id": 2,
          "alertFields": [
            "usage_guest",
            "usage_guest_nice",
            "usage_idle",
            "usage_iowait",
            "usage_irq",
            "usage_nice",
            "usage_softirq",
            "usage_steal",
            "usage_system",
            "usage_user"
          ],
          "alertSubCluster": [
            {
              "_id": 1,
              "alertTypes": [
                {
                  "type": "cpu-total"
                },
                {
                  "type": "cpu0"
                },
                {
                  "type": "cpu1"
                },
                {
                  "type": "cpu10"
                },
                {
                  "type": "cpu11"
                },
                {
                  "type": "cpu12"
                },
                {
                  "type": "cpu13"
                },
                {
                  "type": "cpu14"
                },
                {
                  "type": "cpu15"
                },
                {
                  "type": "cpu16"
                },
                {
                  "type": "cpu17"
                },
                {
                  "type": "cpu18"
                },
                {
                  "type": "cpu19"
                },
                {
                  "type": "cpu2"
                },
                {
                  "type": "cpu20"
                },
                {
                  "type": "cpu21"
                },
                {
                  "type": "cpu22"
                },
                {
                  "type": "cpu23"
                },
                {
                  "type": "cpu24"
                },
                {
                  "type": "cpu25"
                },
                {
                  "type": "cpu26"
                },
                {
                  "type": "cpu27"
                },
                {
                  "type": "cpu28"
                },
                {
                  "type": "cpu29"
                },
                {
                  "type": "cpu3"
                },
                {
                  "type": "cpu30"
                },
                {
                  "type": "cpu31"
                },
                {
                  "type": "cpu32"
                },
                {
                  "type": "cpu33"
                },
                {
                  "type": "cpu34"
                },
                {
                  "type": "cpu35"
                },
                {
                  "type": "cpu36"
                },
                {
                  "type": "cpu37"
                },
                {
                  "type": "cpu38"
                },
                {
                  "type": "cpu39"
                },
                {
                  "type": "cpu4"
                },
                {
                  "type": "cpu40"
                },
                {
                  "type": "cpu41"
                },
                {
                  "type": "cpu42"
                },
                {
                  "type": "cpu43"
                },
                {
                  "type": "cpu44"
                },
                {
                  "type": "cpu45"
                },
                {
                  "type": "cpu46"
                },
                {
                  "type": "cpu47"
                },
                {
                  "type": "cpu5"
                },
                {
                  "type": "cpu6"
                },
                {
                  "type": "cpu7"
                },
                {
                  "type": "cpu8"
                },
                {
                  "type": "cpu9"
                }
              ],
              "name": "cpu"
            },
            {
              "_id": 2,
              "alertTypes": [
                {
                  "type": "2030045418"
                }
              ],
              "name": "host"
            }
          ],
          "name": "cpu"
        },
        {
          "_id": 3,
          "alertFields": [
            "free",
            "inodes_free",
            "inodes_total",
            "inodes_used",
            "total",
            "used",
            "used_percent"
          ],
          "alertSubCluster": [
            {
              "_id": 1,
              "alertTypes": [
                {
                  "type": "sdb1"
                },
                {
                  "type": "sdb2"
                },
                {
                  "type": "sdb5"
                }
              ],
              "name": "device"
            },
            {
              "_id": 2,
              "alertTypes": [
                {
                  "type": "ext4"
                }
              ],
              "name": "fstype"
            },
            {
              "_id": 3,
              "alertTypes": [
                {
                  "type": "2030045418"
                }
              ],
              "name": "host"
            },
            {
              "_id": 4,
              "alertTypes": [
                {
                  "type": "rw"
                }
              ],
              "name": "mode"
            },
            {
              "_id": 5,
              "alertTypes": [
                {
                  "type": "/"
                },
                {
                  "type": "/home"
                },
                {
                  "type": "/usr"
                }
              ],
              "name": "path"
            }
          ],
          "name": "disk"
        },
        {
          "_id": 4,
          "alertFields": [
            "io_time",
            "iops_in_progress",
            "read_bytes",
            "read_time",
            "reads",
            "weighted_io_time",
            "write_bytes",
            "write_time",
            "writes"
          ],
          "alertSubCluster": [
            {
              "_id": 1,
              "alertTypes": [
                {
                  "type": "2030045418"
                }
              ],
              "name": "host"
            },
            {
              "_id": 2,
              "alertTypes": [
                {
                  "type": "loop0"
                },
                {
                  "type": "loop1"
                },
                {
                  "type": "loop2"
                },
                {
                  "type": "nvme0n1"
                },
                {
                  "type": "nvme1n1"
                },
                {
                  "type": "nvme2n1"
                },
                {
                  "type": "nvme3n1"
                },
                {
                  "type": "sda"
                },
                {
                  "type": "sda1"
                },
                {
                  "type": "sda2"
                },
                {
                  "type": "sdb"
                },
                {
                  "type": "sdb1"
                },
                {
                  "type": "sdb2"
                },
                {
                  "type": "sdb3"
                },
                {
                  "type": "sdb5"
                }
              ],
              "name": "name"
            }
          ],
          "name": "diskio"
        },
        {
          "_id": 5,
          "alertFields": [
            "boot_time",
            "context_switches",
            "entropy_avail",
            "interrupts",
            "processes_forked"
          ],
          "alertSubCluster": [
            {
              "_id": 1,
              "alertTypes": [
                {
                  "type": "2030045418"
                }
              ],
              "name": "host"
            }
          ],
          "name": "kernel"
        },
        {
          "_id": 6,
          "alertFields": [
            "active",
            "available",
            "available_percent",
            "buffered",
            "cached",
            "commit_limit",
            "committed_as",
            "dirty",
            "free",
            "high_free",
            "high_total",
            "huge_page_size",
            "huge_pages_free",
            "huge_pages_total",
            "inactive",
            "low_free",
            "low_total",
            "mapped",
            "page_tables",
            "shared",
            "slab",
            "swap_cached",
            "swap_free",
            "swap_total",
            "total",
            "used",
            "used_percent",
            "vmalloc_chunk",
            "vmalloc_total",
            "vmalloc_used",
            "wired",
            "write_back",
            "write_back_tmp"
          ],
          "alertSubCluster": [
            {
              "_id": 1,
              "alertTypes": [
                {
                  "type": "2030045418"
                }
              ],
              "name": "host"
            }
          ],
          "name": "mem"
        },
        {
          "_id": 7,
          "alertFields": [
            "blocked",
            "dead",
            "idle",
            "paging",
            "running",
            "sleeping",
            "stopped",
            "total",
            "total_threads",
            "unknown",
            "zombies"
          ],
          "alertSubCluster": [
            {
              "_id": 1,
              "alertTypes": [
                {
                  "type": "2030045418"
                }
              ],
              "name": "host"
            }
          ],
          "name": "processes"
        },
        {
          "_id": 8,
          "alertFields": [
            "free",
            "in",
            "out",
            "total",
            "used",
            "used_percent"
          ],
          "alertSubCluster": [
            {
              "_id": 1,
              "alertTypes": [
                {
                  "type": "2030045418"
                }
              ],
              "name": "host"
            }
          ],
          "name": "swap"
        },
        {
          "_id": 9,
          "alertFields": [
            "load1",
            "load15",
            "load5",
            "n_cpus",
            "n_users",
            "uptime",
            "uptime_format"
          ],
          "alertSubCluster": [
            {
              "_id": 1,
              "alertTypes": [
                {
                  "type": "2030045418"
                }
              ],
              "name": "host"
            }
          ],
          "name": "system"
        }
      ]
    }

    response = mock.onGet('/api/v1.0/get_alert_types/').reply(200, data);
  });


  test('renders alert table component', async () => {
    const mock = new MockAdapter(axios);
    jest.setTimeout(30000);

    const { getByLabelText, queryAllByText, getByTitle, getAllByTestId, getAllByTitle, getByTestId, getByText, getAllByText, asFragment } = wrapper;

    getByTestId("AlertsTableTag");
    await act(async () => {
      select_input = await waitForElement(() => getAllByTitle('Edit')[0]);
      fireEvent.click(select_input);

      select_input = await waitForElement(() => getAllByTestId("SelectEditTag")[0]);
      fireEvent.click(select_input);

      // select_input = await waitForElement(() => getAllByTestId("SelectEditMenuItemTag")[0]);
      // fireEvent.click(select_input);

      // select_input = await waitForElement(() => getAllByTestId("SelectEditMenuItemTag")[0]);
      // fireEvent.change(select_input, { target: { value: "Less Than" } });
      let spy = jest.spyOn(axios, "post").mockReturnValue(200);

      fireEvent.click(getAllByTitle('Save')[0]);
      await new Promise(resolve => setTimeout(resolve, 1000));
      expect(spy).toBeCalled();//toHave('/api/v1.0/update_alerts/', { "_id":{"$oid":"5d5e67b7f082d4c68c98f58b"},"alertName": "NewAlert", "alertCluster": "cpu", "alertSubCluster": "device", "alertType": "cpu-total", "alertCondition": "Greater Than", "alertField": "usage_system", "description": "last", "alertRange": "109", "active": true, "selected": false, "edit": false },{"headers": {"Accept": "application/json", "Content-Type": "application/json", "x-access-token": null}});

      select_input = await waitForElement(() => getAllByTitle('Edit')[0]);
      fireEvent.click(select_input);

      select_input_child = await waitForElement(() => getAllByTestId("SelectEditTag")[0]);
      fireEvent.click(select_input_child);
      spy = jest.spyOn(axios, "post").mockReturnValue(200);

      fireEvent.click(getAllByTitle('Save')[0]);
      await new Promise(resolve => setTimeout(resolve, 1000));
      expect(spy).toBeCalled();//toHave('/api/v1.0/update_alerts/', { "_id":{"$oid":"5d5e67b7f082d4c68c98f58b"},"alertName": "NewAlert", "alertCluster": "cpu", "alertSubCluster": "device", "alertType": "cpu-total", "alertCondition": "Greater Than", "alertField": "usage_system", "description": "last", "alertRange": "109", "active": true, "selected": false, "edit": false },{"headers": {"Accept": "application/json", "Content-Type": "application/json", "x-access-token": null}});

      select_delete = await waitForElement(() => getAllByTitle('Delete')[0]);
      fireEvent.click(select_delete);

      spy = jest.spyOn(axios, "post").mockReturnValue(200);
      //mock.onPost().reply(200);
      fireEvent.click(getAllByTitle('Save')[0]);
      await new Promise(resolve => setTimeout(resolve, 1000));
      expect(spy).toBeCalled();//toHave('/api/v1.0/update_alerts/', { "_id":{"$oid":"5d5e67b7f082d4c68c98f58b"},"alertName": "NewAlert", "alertCluster": "cpu", "alertSubCluster": "device", "alertType": "cpu-total", "alertCondition": "Greater Than", "alertField": "usage_system", "description": "last", "alertRange": "109", "active": true, "selected": false, "edit": false },{"headers": {"Accept": "application/json", "Content-Type": "application/json", "x-access-token": null}});

      fireEvent.click(getAllByTitle('Edit')[0]);
      spy = jest.spyOn(axios, "post");

      fireEvent.click(getAllByTitle('Save')[0]);
      await new Promise(resolve => setTimeout(resolve, 1000));
      expect(spy).toBeCalled();
    });

    const input = getByTestId("AlertsTableTag").querySelector('button')
    fireEvent.click(input);

    fireEvent.click(getAllByTestId("SwitchTag")[0].querySelector('input'));
    mock.onPost('/api/v1.0/toggle_alert_status/', { alertName: "ThisisAwesome", status: false }).reply({ "status": 200 });
    fireEvent.change(getAllByTestId("SwitchTag")[0].querySelector('input'), { target: { checked: true, value: "Samsung rocks #/" } });
    mock.onPost().reply({ "status": 200 });
    select_save = getAllByTitle('Save')[0];
    select_switch_tag = getAllByTestId("SwitchTag")[0];
    await act(async () => {
      fireEvent.click(select_input);
      //let spy = jest.spyOn(axios, "post");
      mock.onPost('/api/v1.0/update_alerts/').reply({ "status": 200 });
      fireEvent.click(select_save);
      await new Promise(resolve => setTimeout(resolve, 1000));
      //expect(spy).toBeCalled();
    });
  });

  test('renders alert table component status 200', async () => {
    let mock = new MockAdapter(axios);
    jest.setTimeout(30000);
    const { getByLabelText, queryAllByText, getByTitle, getAllByTestId, getAllByTitle, getByTestId, getByText, getAllByText, asFragment } = wrapper;
    //getByTestId("AlertManagementTag");
    getByTestId("AlertsTableTag");
    await act(async () => {
      //select_input = await waitForElement(() => getAllByTitle('Edit')[0]);
      fireEvent.click(select_input);

      // select_input = await waitForElement(() => getAllByTestId("SelectEditTag")[0]);
      // fireEvent.click(select_input);

      // select_input = await waitForElement(() => getAllByTestId("SelectEditMenuItemTag")[0]);
      // fireEvent.click(select_input);

      // select_input = await waitForElement(() => getAllByTestId("SelectEditMenuItemTag")[0]);
      // fireEvent.change(select_input, { target: { value: "Less Than" } });
      let spy = jest.spyOn(axios, "post");
      mock.onPost().reply(200);
      fireEvent.click(select_save);
      await new Promise(resolve => setTimeout(resolve, 1000));
      expect(spy).toBeCalled();//toHave('/api/v1.0/update_alerts/', { "_id":{"$oid":"5d5e67b7f082d4c68c98f58b"},"alertName": "NewAlert", "alertCluster": "cpu", "alertSubCluster": "device", "alertType": "cpu-total", "alertCondition": "Greater Than", "alertField": "usage_system", "description": "last", "alertRange": "109", "active": true, "selected": false, "edit": false },{"headers": {"Accept": "application/json", "Content-Type": "application/json", "x-access-token": null}});

      //select_input = await waitForElement(() => getAllByTitle('Edit')[0]);
      fireEvent.click(select_input);

      //let select_input_child = await waitForElement(() => getAllByTestId("SelectEditTag")[0]);
      fireEvent.click(select_input_child);
      spy = jest.spyOn(axios, "post");

      fireEvent.click(select_save);
      await new Promise(resolve => setTimeout(resolve, 1000));
      expect(spy).toBeCalled();//toHave('/api/v1.0/update_alerts/', { "_id":{"$oid":"5d5e67b7f082d4c68c98f58b"},"alertName": "NewAlert", "alertCluster": "cpu", "alertSubCluster": "device", "alertType": "cpu-total", "alertCondition": "Greater Than", "alertField": "usage_system", "description": "last", "alertRange": "109", "active": true, "selected": false, "edit": false },{"headers": {"Accept": "application/json", "Content-Type": "application/json", "x-access-token": null}});

      //select_delete = await waitForElement(() => getAllByTitle('Delete')[0]);
      fireEvent.click(select_delete);

      spy = jest.spyOn(axios, "post");
      mock.onPost().reply(200);
      fireEvent.click(select_save);
      await new Promise(resolve => setTimeout(resolve, 1000));
      expect(spy).toBeCalled();//toHave('/api/v1.0/update_alerts/', { "_id":{"$oid":"5d5e67b7f082d4c68c98f58b"},"alertName": "NewAlert", "alertCluster": "cpu", "alertSubCluster": "device", "alertType": "cpu-total", "alertCondition": "Greater Than", "alertField": "usage_system", "description": "last", "alertRange": "109", "active": true, "selected": false, "edit": false },{"headers": {"Accept": "application/json", "Content-Type": "application/json", "x-access-token": null}});

      fireEvent.click(select_input);
      spy = jest.spyOn(axios, "post");

      fireEvent.click(select_save);
      await new Promise(resolve => setTimeout(resolve, 1000));
      expect(spy).toBeCalled();
    });

    const input = getByTestId("AlertsTableTag").querySelector('button')
    fireEvent.click(input);

    fireEvent.click(select_switch_tag.querySelector('input'));
    mock.onPost('/api/v1.0/toggle_alert_status/', { alertName: "ThisisAwesome", status: false }).reply(200);
    fireEvent.change(select_switch_tag.querySelector('input'), { target: { checked: true, value: "Samsung rocks #/" } });
    mock.onPost().reply(200);

  });

  
  test('renders alert types component', async () => {
    let mock = new MockAdapter(axios);
    jest.setTimeout(30000);

    const { getByLabelText, queryAllByText, getAllByTestId, getAllByRole, getByRole, getByTestId, getByText, getAllByText, asFragment } = wrapper;
    getByTestId("AlertsTypesTag");
    getAllByTestId("ClusterTag")[0];
    let alert_input;
    let spy;
    let yesBtn;
    await act(async () => {
      //const tree_input =  await waitForElement(() => getAllByRole("treeitem")[0]);
      const tree_input = await waitForElement(() => getAllByRole("treeitem")[0].querySelector('div'));
      fireEvent.click(tree_input);
      const subtreeItem = await waitForElement(() => getAllByRole("subtreeitem")[0].querySelector('div'));
      fireEvent.click(subtreeItem);
      const formControl = await waitForElement(() => getAllByTestId("alertTypesCheckbox")[0]);
      fireEvent.click(formControl);
      fireEvent.change(formControl, { target: { checked: true } });


      // let { getByLabelText, queryAllByText, getAllByTestId, getAllByRole, getByRole, getByTestId, getByText, getAllByText, asFragment } = render(
      //   <I18nextProvider i18n={i18n}>
      //       <Provider store={store}>
      //           <AlertFields alertClusterList={alertClusterList} radioindex={radioindex} alertClusterName={alertClusterName} selectedAlertSubCluster={selectedAlertSubCluster} alertType={alertType} />
      //       </Provider>
      //   </I18nextProvider>);
      //await waitForElement(() => getByLabelText(/usage_idle/i));
      // const input =  await waitForElement(() => getByRole(/radiogroup/i));
      const input = await waitForElement(() => getAllByTestId(/alertFieldRadioTag/i)[0]);
      fireEvent.click(input);

      alert_input = await waitForElement(() => getByLabelText(/Alert Name/i));
      fireEvent.click(alert_input);
      fireEvent.change(alert_input, { target: { value: "Alert" } });
      fireEvent.keyDown(alert_input, { key: 'Enter', keyCode: 13, charCode: 13 });


      let dropdown_input = await waitForElement(() => getAllByTestId("selectAddNewAlertsTag")[0]);
      fireEvent.click(dropdown_input);

      // let dropdown_element = await waitForElement(() => getAllByTestId("selectMenuItemAddNewAlertsTag")[0]);
      // fireEvent.click(dropdown_element);

      // fireEvent.change(dropdown_element, { target: { value: "Less Than" } });

      //let value_input = await waitForElement(() => getByLabelText(/Value/i));
      let value_input = await waitForElement(() => getByTestId(/Alert_Range_TextField/i));
      fireEvent.click(value_input);
      fireEvent.change(value_input, { target: { value: "60" } });
      fireEvent.keyDown(value_input, { key: 'Enter', keyCode: 13, charCode: 13 });

      let button_input = await waitForElement(() => getAllByTestId('ButtonTag')[0].querySelector('button'));
      spy = jest.spyOn(axios, "post");
      mock.onPost('/api/v1.0/add_alert/').reply(200);
      fireEvent.click(button_input);
      yesBtn = await waitForElement(() => getByText('Yes'));
      spy = jest.spyOn(axios, "post");
      mock.onPost('/api/v1.0/add_alert/').reply(200);
      fireEvent.click(yesBtn);
      expect(spy).toBeCalled();

    });

    fireEvent.click(getByTestId("AlertsTypesTag").querySelector('ul'));

  });



  test('renders add new alert component status 200', async () => {
    let mock = new MockAdapter(axios);
    jest.setTimeout(30000);
    const { getByLabelText, queryAllByText, getByTitle, getAllByTestId, getAllByTitle, getByTestId, getByText, getAllByText, asFragment } = wrapper;
    getAllByTestId("addNewAlertsTag")[0];
    let alert_input;
    let spy;
    let yesBtn;
    await act(async () => {
      alert_input = await waitForElement(() => getByLabelText(/Alert Name/i));
      fireEvent.click(alert_input);
      fireEvent.change(alert_input, { target: { value: "Alert" } });
      fireEvent.keyDown(alert_input, { key: 'Enter', keyCode: 13, charCode: 13 });
      // alert_input = await waitForElement(() => getAllByTestId('ButtonTag')[0].querySelector('button'));
      // fireEvent.click(alert_input);
      // yesBtn = await waitForElement(() => getByText('Yes'));
      // spy = jest.spyOn(axios, "post");
      // mock.onPost('/api/v1.0/add_alert/').reply(200);
      // fireEvent.click(yesBtn);
      // expect(spy).toBeCalled();

      // alert_input = await waitForElement(() => getByLabelText(/Alert Name/i));
      // fireEvent.click(alert_input);
      // fireEvent.change(alert_input, { target: { value: "AlertName" } });
      // fireEvent.keyDown(alert_input, { key: 'Enter', keyCode: 13, charCode: 13 });

      // alert_input = await waitForElement(() => getAllByTestId("alertFieldRadioTag")[0]);
      // fireEvent.click(alert_input);

      let field_input = await waitForElement(() => getByLabelText(/Alert Field/i));
      fireEvent.click(field_input);
      fireEvent.change(field_input, { target: { value: "Usage_system" } });

      let dropdown_input = await waitForElement(() => getAllByTestId("selectAddNewAlertsTag")[0]);
      fireEvent.click(dropdown_input);

      // let dropdown_element = await waitForElement(() => getAllByTestId("selectMenuItemAddNewAlertsTag")[0]);
      // fireEvent.click(dropdown_element);

      // //let dropdown_element2 = await waitForElement(() => getAllByTestId("selectMenuItemAddNewAlertsTag")[0]);
      // fireEvent.change(dropdown_element, { target: { value: "Less Than" } });

      //let value_input = await waitForElement(() => getByLabelText(/Value/i));
      let value_input = await waitForElement(() => getByTestId(/Alert_Range_TextField/i));
      fireEvent.click(value_input);
      fireEvent.change(value_input, { target: { value: "60" } });
      fireEvent.keyDown(value_input, { key: 'Enter', keyCode: 13, charCode: 13 });

      let button_input = await waitForElement(() => getAllByTestId('ButtonTag')[0].querySelector('button'));
      spy = jest.spyOn(axios, "post");
      mock.onPost('/api/v1.0/add_alert/').reply(200);
      fireEvent.click(button_input);
      yesBtn = await waitForElement(() => getByText('Yes'));
      spy = jest.spyOn(axios, "post");
      mock.onPost('/api/v1.0/add_alert/').reply(200);
      fireEvent.click(yesBtn);
      expect(spy).toBeCalled();
    });
  });

  test('renders alert management container', () => {
    const { getByTestId, asFragment } = wrapper;
    expect(asFragment()).toMatchSnapshot();
    getByTestId('alertManagementTag')
  });


});