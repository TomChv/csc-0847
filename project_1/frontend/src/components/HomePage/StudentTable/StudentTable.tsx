import React, {useEffect, useState} from "react";
import {
    IconButton,
    Paper,
    Table,
    TableBody,
    TableCell,
    TableContainer,
    TableHead,
    TableRow
} from "@mui/material";
import {users} from "../../../mocks/users";
import StudentRow from "./StudentRow";
import SearchIcon from '@mui/icons-material/Search';
import AddIcon from '@mui/icons-material/Add';
import {client} from "../../../backend/client";
import AddStudentForm from "../AddStudentForm/AddStudentForm";
import SearchStudentForm from "../SearchStudentForm/SearchStudentForm";

export default function StudentTable() {
    const [usersState, setUsersState] = useState(users);
    const [displayAddForm, setDisplayAddForm] = useState(false);
    const [displaySearchForm, setDisplaySearchForm] = useState(false);
    const [update, setUpdate] = useState(false);

    useEffect(() => {
        console.log("Fetch users...")
        client.listUsers()
            .then((users) => setUsersState(users))
            .catch((e) => console.error(`could not fetch users ${e.message}`))
    }, [displayAddForm, update])

    const searchUser = () => {
        setDisplaySearchForm(true)
    }

    const addUser = () => {
        setDisplayAddForm(true)
    }

    return (
        <TableContainer component={Paper} sx={{maxWidth: 1000}}>
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
                            <IconButton onClick={searchUser}>
                                <SearchIcon/>
                            </IconButton>
                        </TableCell>
                        <TableCell align="left">
                            <IconButton onClick={addUser}>
                                <AddIcon/>
                            </IconButton>
                        </TableCell>
                    </TableRow>
                </TableHead>
                <TableBody>
                    {usersState.map((user) => (<StudentRow
                        user={user}
                        update={update}
                        setUpdate={setUpdate}
                        key={user.id}
                    />))}
                </TableBody>
            </Table>

            <AddStudentForm open={displayAddForm} setOpen={setDisplayAddForm}/>
            <SearchStudentForm
                users={usersState} setUsers={setUsersState}
                open={displaySearchForm} setOpen={setDisplaySearchForm}
            />
        </TableContainer>
    )
}