import React, { useState, useEffect } from 'react';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';
import DeleteIcon from '@material-ui/icons/Delete';
import { purple } from '@material-ui/core/colors';

import {
  Content,
  Header,
  Page,
  pageTheme,
} from '@backstage/core';
import { createTheme, lightTheme, darkTheme } from '@backstage/theme';
import PersonAddRoundedIcon from '@material-ui/icons/PersonAddRounded';
import HomeRoundedIcon from '@material-ui/icons/HomeRounded';
import { EntProvince } from '../../api';


const useStyles = makeStyles((theme: Theme) =>
 createStyles({
  table: {
    minWidth: 650,
  },
  buttonRight: {
    marginLeft: theme.spacing(150),
    marginBottom: theme.spacing(2),
  },
  }),
);
 
export default function ProvinceTable() {
  const classes = useStyles();
  const http = new DefaultApi();
  const [provinces, setProvince] = useState<EntProvince[]>([]);
  const [loading, setLoading] = useState(true);
  

  // get province
  useEffect(() => {
    const getProvince = async () => {
      const res = await http.listProvince({ limit: 50, offset: 0 });
      setLoading(true);
      setProvince(res);
      console.log(res);
    };
    getProvince();
  }, [loading]);
  
 
  // ui 
 return (
  <Page theme={pageTheme.tool}>
    <Header title={`ระบบจัดเก็บข้อมูลจังหวัด`}>
    <div>&nbsp;&nbsp;&nbsp;</div>
    <Button style={{ marginLeft: 100, backgroundColor: 'white', width: '150px' }} href="/" startIcon={<HomeRoundedIcon/>}> home</Button>
    </Header>
    
    <Content>
   <TableContainer component={Paper}>
     <Table className={classes.table} aria-label="simple table">
       <TableHead>
         <TableRow>
           <TableCell align="center">No</TableCell>
           <TableCell align="center">Student</TableCell>
           <TableCell align="center">Province</TableCell>
           <TableCell align="center">District</TableCell>
           <TableCell align="center">Subdistrict</TableCell>
         </TableRow>
       </TableHead>

       <TableBody>
       {provinces.map((item:any) => (
           <TableRow key={item.id}>
             <TableCell align="center">{item.id}</TableCell>
             <TableCell align="center">{item.edges?.student?.student}</TableCell>
             <TableCell align="center">{item.name}</TableCell>

             <TableCell align="center">{item.edges?.district?.district}</TableCell>
             <TableCell align="center">{item.edges?.subdistrict?.subdistrict}</TableCell>
             <TableCell align="center">
     
             </TableCell>

           </TableRow>
         ))}
       </TableBody>
     </Table>
   </TableContainer>
   </Content>
  </Page>
);
}