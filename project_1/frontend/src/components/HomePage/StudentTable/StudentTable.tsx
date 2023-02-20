import React from "react";
import {Box, Paper, Table, TableBody, TableCell, TableContainer, TableHead, TableRow} from "@mui/material";
import {users} from "../../../mocks/users";
import StudentRow from "./StudentRow";
import SearchIcon from '@mui/icons-material/Search';
import AddIcon from '@mui/icons-material/Add';

export default function StudentTable() {
    const usersData = users;

    const searchUser = () => {
        console.log("Search an user")
    }

    const addUser = () => {
        console.log("addUser")
    }

    return (
        <TableContainer component={Paper} sx={{ maxWidth: 1000 }}>
            <Table sx={{minWidth: 700}} aria-label="student table">
                <TableHead>
                    <TableRow>
                        <TableCell>Student ID</TableCell>
                        <TableCell align="left">Firstname</TableCell>
                        <TableCell align="left">Lastname</TableCell>
                        <TableCell align="left">Email</TableCell>
                        <TableCell align="left">Mailing Address</TableCell>
                        <TableCell align="left">GPA</TableCell>
                        <TableCell align="left">
                            <Box onClick={searchUser}>
                                <SearchIcon/>
                            </Box>
                        </TableCell>
                        <TableCell align="left">
                            <Box onClick={addUser}>
                                <AddIcon/>
                            </Box>
                        </TableCell>
                    </TableRow>
                </TableHead>
                <TableBody>
                    {usersData.map((user) => (<StudentRow user={user} key={user.id}/>))}
                </TableBody>
            </Table>
        </TableContainer>
    )
}