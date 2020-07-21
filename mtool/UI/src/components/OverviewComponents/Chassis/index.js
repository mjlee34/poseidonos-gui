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

DESCRIPTION: Overview Page Chassis Component
@NAME : index.js
@AUTHORS: Jay Hitesh Sanghavi
@Version : 1.0 
@REVISION HISTORY
[03/11/2019] [Jay] : Prototyping..........////////////////////
*/

import React, { Component } from 'react';
import { connect } from 'react-redux';
import { withStyles } from '@material-ui/core/styles';
import MaterialTable from 'material-table';
import { Paper, GridList, Typography, Tooltip, GridListTile , createMuiTheme,InputLabel, } from '@material-ui/core';
import Grid from '@material-ui/core/Grid';

import ThemeProvider from '@material-ui/core/styles/MuiThemeProvider';
import Search from '@material-ui/icons/Search';
import SaveAlt from '@material-ui/icons/SaveAlt';
import ChevronLeft from '@material-ui/icons/ChevronLeft';
import ChevronRight from '@material-ui/icons/ChevronRight';
import FirstPage from '@material-ui/icons/FirstPage';
import LastPage from '@material-ui/icons/LastPage';
import Add from '@material-ui/icons/Add';
import Check from '@material-ui/icons/Check';
import FilterList from '@material-ui/icons/FilterList';
import Remove from '@material-ui/icons/Remove';
import EditIcon from '@material-ui/icons/Edit';
import TrashIcon from '@material-ui/icons/Delete';
import ArrowUpward from '@material-ui/icons/ArrowUpward';
import Clear from '@material-ui/icons/Clear';
import './Chassis.css';
import Legend from '../../Legend'
import * as actionTypes from '../../../store/actions/actionTypes';
import DiskDetails from '../../DiskDetails';
import ServerInformation from "../ServerInformation"

const styles = theme => ({
  multilineColor: {
    color: 'black'
  },
  root: {
    flexGrow: 1,
  },

  overviewPaper: {
    marginTop: theme.spacing(2),
    width: '100%',
  },


  chassisOuterGrid: {
   // border: '1px solid gray',
    maxWidth: '100%',
    flexBasis: '100%',
    overflowY: 'auto',
    overflowX: 'hidden',
    background: '#fff',
  },
  chassisInnerGrid: {
    maxWidth: '100%',
  },

  label: {
    marginLeft: theme.spacing(1),
    marginTop: theme.spacing(3),
  },
  gridTile: {
    width: 450,
    minWidth: 10,
    border: '2px solid lightgray',
    display: 'flex',
    justifyContent: 'center',
    alignItems: 'center',
    position: 'relative',
    flexDirection: 'column',
    '&>div': {
      height: 'auto'
    }
  },
  gridTileDisabled: {
    backgroundColor: '#e2e1e1'
  },
  gridTileHealthy: {
    backgroundColor: 'green'
  },
  chassisGridContainer: {
    width: '100%',
    overflowX: 'auto',
    [theme.breakpoints.down('xs')]: {
      width: 'calc(100% - 32px)'
    }
  },
  chassisContainer: {
    margin: '20px',
    padding: '5px',
    border: '1px solid gray',

  },
  diskContainer: {
    display: 'flex',
    flexWrap: 'wrap',
    justifyContent: 'space-around',
    overflow: 'hidden',
    backgroundColor: theme.palette.background.paper,
    padding: theme.spacing(0, 2, 0, 2),
    minWidth: 120
  },
  gridList: {
    flexWrap: 'nowrap',
    // Promote the list into his own layer on Chrome. This cost memory but helps keeping high FPS.
    transform: 'translateZ(0)',
    flexGrow: 1,
    padding: theme.spacing(0, 0)
  },

  legendContainer: {
    maxWidth: '100%',
    padding: theme.spacing(2, 2)
  },
  chassisLabel:{
    padding:'15px',
    width:'100%',
    textAlign:'center'
  },
  
  powerParentGrid: {
   marginBottom: -theme.spacing(0.5)       
  }
});

const defaultDiskDetails = {
  DevicePath: 'NA',
  BuildDate: 'NA',
  Manufacturer: 'NA',
  PartNumber: 'NA',
  SerialNumber: 'NA',
  Model: 'NA',
  PredictedMediaLifeLeftPercent:'NA',
  PhysicalSize: 'NA',
  UsedBytes: 'NA',
  Firmware: 'NA',
  critical_warning: 'NA',
  temperature: 'NA',
  avail_spare: 'NA',
  spare_thresh: 'NA',
  precent_used: 'NA',
  data_units_read: 'NA',
  data_units_written: 'NA',
  critical_comp_time: 'NA',
  warning_temp_time: 'NA',
  percent_used: 'NA',
};

class Chassis extends Component {
  constructor(props) {
    super(props);
    this.showPopup = this.showPopup.bind(this);
    this.closePopup = this.closePopup.bind(this);
    this.getDiskDetails = this.getDiskDetails.bind(this);
    this.theme = createMuiTheme({
      overrides: {
        MuiSvgIcon: {
          //  stylesheet name
          root: {
            //  rule name
            color: '#808080',
          },
        },
        MuiTablePagination: {
          menuItem: {
            fontSize: "12px",
            minHeight: "0px"
          },
          select: {
            width: "45px"
          },
          toolbar: {
            maxHeight: '50px',
            minHeight: '20px'
          },
        },
      },

      palette: {
        primary: {
          main: '#4caf50',
        },
        secondary: {
          main: '#808080',
        },
      },

    });

    this.state = {
      disknameandslot:'',
      diskDetails: {...defaultDiskDetails},
      popupOpen: false,
      columns: [
        {
          title: 'Name',
          field: 'chassisname',
          cellStyle: {
            fontFamily:'Arial',
            fontSize: '12px',
          },
        },
        {
          title: 'Status',
          field: 'chassisstatus',
          cellStyle: {
            fontFamily:'Arial',
            fontSize: '12px',
          },
        },
        {
          title: 'Serial',
          field: 'chassisserial',
          cellStyle: {
            fontFamily:'Arial',
            fontSize: '12px',
          },
        },
        {
          title: 'Other',
          field: 'chassisother',
          cellStyle: {
            fontFamily:'Arial',
            fontSize: '12px',
          },
        },
      ],
      data: [{ 'chassisname': 'ETH0', 'chassisstatus': 'Healthy', 'chassisserial': 'CNGODYTRNHD', 'chassisother': 'SPEED Slot 40GB' }, { 'chassisname': 'ETH1', 'chassisstatus': 'Healthy', 'chassisserial': 'CNGODYTRNHDH', 'chassisother': 'SPEED Slot 50GB' },{ 'chassisname': 'PSU1', 'chassisstatus': 'Healthy', 'chassisserial': 'POWERCNGODYTRNHDH', 'chassisother': ' - ' }]
    };
  }

  getDiskDetails(name,slot,disk) {
    this.setState({
      ...this.state,
      diskDetails:
      {
        ...defaultDiskDetails,
        ...disk
      },
      disknameandslot: `Disk Details (Disk Name: ${ name }${slot }, Slot Number: ${ slot+1 })`,
      popupOpen: true,
    });
  }

  showPopup(name,slot,disk) {
    this.getDiskDetails(name,slot,disk)
  }

  closePopup() {
    this.setState({
      ...this.state,
      popupOpen: false,
    });
  }

  componentDidMount() {
    this.props.fetchChassisFrontInfo();
    // this.props.fetchChassisRearInfo();
  }

  render() {
    const { classes } = this.props;
    const freeSlots = [];
    if (this.props.chassis_front_list) {
      for (let i = this.props.chassis_front_list.length; i <32; i += 1) {
        freeSlots.push(
          <Grid className={`${classes.gridTile} ${classes.gridTileDisabled}`}>
            <Typography color="secondary" className={classes.diskNo}>{i + 1}</Typography>
          </Grid>
        );
      }
    }

    return (
      // <Paper className={classes.overviewPaper}>
        <Grid item spacing={2} container className={classes.powerParentGrid} data-testid="Chassis_Component">
          <Grid sm={6} xs={12} item container>
            <Paper xs={12} item className={classes.chassisOuterGrid}>  
            <InputLabel className={classes.chassisLabel}>Chassis Front View</InputLabel>
              <Grid sm={6} xs={12} item className={classes.chassisInnerGrid}>

                <div className={classes.chassisGridContainer}>
                  {/* <Grid container className={classes.chassisContainer}>
                <GridListTile></GridListTile> */}

                  <Grid container className={classes.diskContainer}>
                    {/* <InputLabel className={classes.chassisLabel}>Chassis</InputLabel> */}
                    <GridList cellHeight={150} className={classes.gridList} cols={32} data-testid="ChassisDiskPopUp">
                      {this.props.chassis_front_list
                        ? this.props.chassis_front_list.map((disk, index) => {
                          return (
                            <Tooltip
                            data-testid="Tooltip"
                              classes={{
                                tooltip: classes.tooltip,
                              }}
                              title={(
                                <React.Fragment>
                                  <div>
                                    Name:
                            {`nvme${ index}`}
                                  </div>
                                  {/* <div>
                                    Size:
                            {formatBytes(disk.size * 4 * 1024)}
                                  </div> */}
                                  
                                  <div
                                  data-testid="POPUP"
                                    onClick={() => this.showPopup('nvme',index,disk)}
                                    style={{
                                      cursor: 'pointer',
                                      textAlign: 'right',
                                      margin: '10px',
                                    }}
                                  >
                                    <u>More Details</u>
                                  </div>
                                </React.Fragment>
                              )}
                              interactive
                            >
                              <Grid className={`${classes.gridTile} ${classes.gridTileHealthy}`}>
                                <GridListTile

                                  id={index}
                                // onClick={() => {
                                //   this.toggleRowSelect(index, slot);
                                // }}
                                >
                                  <Typography color="secondary">{index + 1}</Typography>
                                </GridListTile>
                              </Grid>
                            </Tooltip>
                          );
                        })
                        : null}
                      {freeSlots}
                    </GridList>
                  </Grid>
                  {/* </Grid> */}
                </div>
                <Grid item container sm={8} xs={12} wrap="wrap" className={classes.legendContainer}>
                  <Legend bgColor="green" title="Healthy" />
                  <Legend bgColor="yellow" title="Warning" />
                  <Legend bgColor="orange" title="Corrupt" />
                  <Legend bgColor="#e2e1e1" title="No Disk" />
                </Grid>
              </Grid>
            </Paper>
          </Grid>
          <Grid sm={6} xs={12} item container>
          <Paper xs={12} item className={classes.chassisOuterGrid}>       
          <ServerInformation />    
          </Paper>
          </Grid>     
          {/* <Grid sm={6} xs={12} item container >
            <Grid xs={12} item className={classes.chassisOuterGrid}>         
          <InputLabel className={classes.chassisLabel}>Chassis Rear View Information</InputLabel>
             
              <ThemeProvider theme={this.theme} >
                <MaterialTable 
                  icons={{
                    Check,
                    FirstPage,
                    LastPage,
                    NextPage: ChevronRight,
                    PreviousPage: ChevronLeft,
                    Search,
                    ThirdStateCheck: Remove,
                    DetailPanel: ChevronRight,
                    Export: SaveAlt,
                    Filter: FilterList,
                    Add,
                    Edit: EditIcon,
                    Delete: TrashIcon,
                    SortArrow: ArrowUpward,
                    Clear,
                  }}
                  columns={this.state.columns}
                  data={this.state.data}
                  options={{
                    actionsColumnIndex: -1,
                    selection: false,
                    sorting: true,
                    toolbar: false,
                    maxBodyHeight: '170px',
                    rowStyle: {
                      fontSize: '4px',
                    },
                    search: false,
                    paginationType: 'normal',
                    loadingType: 'linear',
                    headerStyle: {
                      backgroundColor: '#788595',
                      color: 'rgba(255, 255, 255, 0.87)',
                      fontSize: '14px',
                      height: '10%',
                      paddingTop: '2px',
                      paddingBottom: '2px',
                    },
                  }}

                />
              </ThemeProvider>
            </Grid>
          </Grid> */}
        
        <DiskDetails
          title={this.state.disknameandslot}
          details={this.state.diskDetails}
          open={this.state.popupOpen}
          onConfirm={this.closePopup}
          handleClose={this.closePopup}
          note_msg="Note: SMART Information cannot be retrieved from the disk as NVMe MI functionality is not implemented in BMC."
        />
        </Grid>
      // </Paper>
    );
  }
}


const mapStateToProps = state => {
  return {
    chassis_front_list: state.hardwareOverviewReducer.chassis_front_list,
    chassis_rear_list: state.hardwareOverviewReducer.chassis_rear_list,
  };
}

const mapDispatchToProps = dispatch => {
  return {
    fetchChassisFrontInfo: () => dispatch({ type: actionTypes.SAGA_HARDWARE_OVERVIEW_FETCH_CHASSIS_FRONT_INFORMATION, }),
    fetchChassisRearInfo: () => dispatch({ type: actionTypes.SAGA_HARDWARE_OVERVIEW_FETCH_CHASSIS_REAR_INFORMATION }),
  };
}

export default withStyles(styles)((connect(mapStateToProps, mapDispatchToProps))(((Chassis))));