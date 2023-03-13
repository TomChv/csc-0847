import AppBar from '@mui/material/AppBar';
import {Box, Typography} from "@mui/material";

export default function TopBar() {
    return (
        <Box sx={{flexGrow: 1}}>
            <AppBar position="sticky" sx={{padding: "5px", height: "80px"}}>
                <Box sx={{ margin: "auto" }}>
                    <Typography variant="h5" component="div" sx={{flexGrow: 1}} align={"center"} fontWeight={"fontWeightMedium"}>
                        CSC 0847 - Project 2
                    </Typography>
                </Box>
            </AppBar>
        </Box>
    )
}