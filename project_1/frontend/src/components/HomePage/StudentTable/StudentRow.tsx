import {TableCell, TableRow} from "@mui/material";
import {User} from "../../../types/user";
import React from "react";
import EditIcon from '@mui/icons-material/Edit';
import DeleteIcon from '@mui/icons-material/Delete';

export interface StudentRowParam {
    user: User
}

export default function StudentRow({user}: StudentRowParam) {
    const deleteUser = () => {
        console.log(`Delete ${user.id}`)
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
                <div onClick={editUser}>
                    <EditIcon/>
                </div>
            </TableCell>
            <TableCell align="left">
                <div onClick={deleteUser}>
                    <DeleteIcon/>
                </div>
            </TableCell>
        </TableRow>
    )
}