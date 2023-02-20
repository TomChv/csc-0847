import {Box, Link, Paper, Typography} from "@mui/material";
import React from "react";
import GitHubIcon from '@mui/icons-material/GitHub';


export default function Footer() {
    return (
        <Paper sx={{
            bottom: 0,
            width: "100%",
            position: "fixed",
            backgroundColor: "#F6F6F6",
            height: "60px",
            padding: "5px",
            display: "flex",
            alignItems: "center"
        }}>
            <Box marginLeft={"10px"} display={"flex"} flexDirection={"row"}>
                <Link href="https://github.com/TomChv/csc-0847/tree/main/project_1" target="_blank">
                    <GitHubIcon fontSize={"large"}/>
                </Link>
                <Typography variant="h5" component="div" sx={{flexGrow: 1, marginLeft: "8px", paddingTop: "3px"}}>
                    Made by TomChv
                </Typography>
            </Box>
        </Paper>
    )
}