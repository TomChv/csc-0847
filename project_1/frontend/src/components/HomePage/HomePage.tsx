import React from "react";
import {Box} from "@mui/material";
import StudentTable from "./StudentTable/StudentTable";

export default function HomePage() {
    return (
        <Box display={"flex"} justifyContent={"center"} alignContent={"center"} marginTop={"5vh"}>
            <StudentTable/>
        </Box>
    )
}