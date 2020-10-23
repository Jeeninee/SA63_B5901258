import React, { FC, useEffect, useState } from 'react';
import TextField from '@material-ui/core/TextField';
import { createStyles, makeStyles, Theme } from '@material-ui/core/styles';
import FormControl from '@material-ui/core/FormControl';
import { InputLabel, IconButton, MenuItem,  } from '@material-ui/core';
import Select from '@material-ui/core/Select';
import NativeSelect from '@material-ui/core/NativeSelect';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import { Alert, AlertTitle } from '@material-ui/lab';
import ContentHeader from '@backstage/core';
import { Link as RouterLink } from 'react-router-dom';
import { DefaultApi } from '../../api/apis';
import { EntProvince} from '../../api/models/EntProvince';
import { EntDistrict} from '../../api/models/EntDistrict';
import { EntStudent } from '../../api/models/EntStudent';
import { EntSubdistrict } from '../../api/models/EntSubdistrict';


const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      display: 'flex',
      flexWrap: 'wrap',
      justifyContent:'center',      
      '& > *': {
        margin: theme.spacing(5),
        width: theme.spacing(52),
        height: theme.spacing(65),
      },
      '& .MuiTextField-root': 
      {
        
        minWidth: 250,

       
      },
    },
    formControl: {
      margin: theme.spacing(2),
      minWidth: 250,
      marginRight:theme.spacing(10),
      marginLeft:theme.spacing(10),
    },
    selectEmpty: {
      marginTop: theme.spacing(2),
      justifyContent:'center',
    },

    
  }),
);


interface province {
  name: string;
  district: number;
  student: number;
  subdistrict: number;
  
 };

export default function Province() {
  const classes = useStyles();
  const http = new DefaultApi();

  const [provinces, setProvinces] = React.useState<EntProvince[]>([]);
  const [districts, setDistricts] = React.useState<EntDistrict[]>([]);
  const [students, setStudents] = React.useState<EntStudent[]>([]);
  const [subdistricts, setSubdistricts] = React.useState<EntSubdistrict[]>([]);

  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);
  const [loading, setLoading] = useState(true);


  const [district, setDistrict] = useState(Number);
  const [subdistrict, setSubdistrict] = useState(Number);
  const [student, setStudent] = useState(Number);
  const [name, setName] = useState(String);


  const getProvince = async () =>{
    const res = await http.listProvince({ limit: 5, offset: 0 });
    setProvinces(res);
  }

  const getDistrict = async () => {
    const res = await http.listDistrict({ limit: 4, offset: 0 });
    setDistricts(res);
  };

console.log(districts)
  const getStudent = async () => {
    const res = await http.listStudent({ limit: 11, offset: 0 });
    setStudents(res);
  };

  const getSubdistrict = async () => {
    const res = await http.listSubdistrict({ limit: 5, offset: 0 });
    setSubdistricts(res);
  };



   // Lifecycle Hooks
  useEffect(() => {
    getDistrict();
    getStudent();
    getSubdistrict();
    getProvince();
   }, []);

  const ProvincehandleChange = (event: React.ChangeEvent<{ value: unknown }>) =>  {
    setName(event.target.value as string);
  };

  const DistricthandleChange = (event: React.ChangeEvent<{ value: unknown }>) =>  {
    setDistrict(event.target.value as number);
  };
  
  const StudenthandleChange = (event: React.ChangeEvent<{ value: unknown }>) =>  {
    setStudent(event.target.value as number);
  };
  
  const SubdistrictChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setSubdistrict(event.target.value as number);
  };

  // create province
const CreateProvince = async () => {
  const province = {
    name: name,
    district: district,
    student: student,
    subdistrict: subdistrict,
  };
    console.log(province);
    const res: any = await http.createProvince({province});
    console.log("bruhhhhhhhhh");
    setStatus(true);

    if (res.id != '') {
      setAlert(true);
    } else {
      setAlert(false);
    }
    const timer = setTimeout(() => {
      setStatus(false);
    }, 3000);   
  }

  return (
    
    <form className={classes.root} noValidate autoComplete="on">
     
      <Paper>
      
        <AppBar
          color="secondary"
          position = "sticky"
          >
          <Toolbar>
            <Typography variant="h6"
            >ระบบจัดเก็บรายชิ่อจังหวัด</Typography>
          </Toolbar>
        </AppBar>

        {status ? (
           <div>
             {alert ? (
               <Alert severity="success" style={{marginTop:5 }}>
                 บันทึกข้อมูลสำเร็จ!
               </Alert>
             ) : (
               <Alert severity="warning" style={{width: 20 ,marginLeft:-10,marginTop:5 }}>
                 This is a warning alert — check it out!
               </Alert>
             )}
           </div>
          ) : null}
          
    <div>
    <FormControl variant="outlined" className={classes.formControl} >
        <InputLabel htmlFor="outlined-age-native-simple">ชื่อนักศึกษา</InputLabel>
        <Select
          name="student"
          value={student}
          onChange={StudenthandleChange}
        >
           {
           
           students.map(item => {
              return (
                <MenuItem key={item.id} value={item.id}>{item.student}</MenuItem>
              );
            })}
        </Select>
      </FormControl>
    </div>

    
    <FormControl variant="outlined" className={classes.formControl} >
    <div>
          <TextField
            id="name"
            label="ชื่อจังหวัด"
            value={name}
            variant="outlined"
            onChange={ProvincehandleChange}     
          />
        </div>
    </FormControl>
    <div> 
    <FormControl variant="outlined" className={classes.formControl}>
      <InputLabel htmlFor="outlined-age-native-simple">ชื่ออำเภอ</InputLabel>
      <Select
        name="district"
       value={district}
        onChange={DistricthandleChange}
      >
        {
          districts.map(item => {
            return (
              <MenuItem key={item.id} value={item.id}>{item.district}</MenuItem>
            );
          })}    
      </Select>
    </FormControl>
    </div>

    <div> 
    <FormControl variant="outlined" className={classes.formControl}>
     <InputLabel htmlFor="outlined-age-native-simple">ชื่อตำบล</InputLabel>
       <Select
          name="subdistrict"
          value={subdistrict}
          onChange={SubdistrictChange}
     
        >
      {subdistricts.map(item => {
          return (
            <MenuItem key={item.id} value={item.id}>{item.subdistrict}</MenuItem>
          );
        })}    
     </Select>
    </FormControl>
    </div>


    <div>
    <Button style={{ marginLeft: 100}} onClick={() => {CreateProvince();}} variant="contained" color="secondary"> บันทึกข้อมูล </Button>
    <Button style={{ marginLeft: 10}}component={RouterLink} to="/table" variant="contained" color="primary"> ดูข้อมูล </Button>
    </div>
    </Paper>
    
    
    
  </form>
  );
}