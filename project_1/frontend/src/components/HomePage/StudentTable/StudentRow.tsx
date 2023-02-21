import {IconButton, TableCell, TableRow} from "@mui/material";
import {User} from "../../../types/user";
import React from "react";
import EditIcon from '@mui/icons-material/Edit';
import DeleteIcon from '@mui/icons-material/Delete';
import {client} from "../../../backend/client";

export interface StudentRowParam {
    user: User
    update: boolean
    setUpdate: React.Dispatch<React.SetStateAction<boolean>>
}

export default function StudentRow({user, update, setUpdate}: StudentRowParam) {
    const deleteUser = () => {
        console.log(`Delete ${user.id}`)

        client.deleteUser(user.id)
            .then(() => setUpdate(!update));
    }

    const editUser = () => {
        console.log(`Edit ${user.id}`)
    }

    return (
        <TableRow>
            <TableCell component={"th"} scope={"row"}>{user.student_id}</TableCell>
            <TableCell align="left">{user.firstname}</TableCell>
            <TableCell align="left">{user.lastname}</TableCell>
            <TableCell align="left">{user.email}</TableCell>
            <TableCell align="left">{user.mailing_address}</TableCell>
            <TableCell align="left">{user.gpa}</TableCell>
            <TableCell align="left">
                <IconButton onClick={editUser}>
                    <EditIcon/>
                </IconButton>
            </TableCell>
            <TableCell align="left">
                <IconButton onClick={deleteUser}>
                    <DeleteIcon/>
                </IconButton>
            </TableCell>
        </TableRow>
    )
}