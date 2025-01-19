import React from 'react';
import { Box, TextField } from '@mui/material';


const Loginform: React.FC = () => {
    return (
        <Box
        component="form"
        sx={{ '& .MuiTextField-root': { m: 1, width: '25ch' } }}
        noValidate
        autoComplete="off">
            <div>
            <TextField id="outlined-basic" label="Outlined" variant="outlined" />
            </div>
        </Box>
    )
}

export default Loginform