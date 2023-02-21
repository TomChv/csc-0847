import React, {useState} from "react";
import {User} from "../../../types/user";
import {Button, Dialog, DialogActions, DialogContent, DialogContentText, DialogTitle, TextField} from "@mui/material";

interface SearchStudentFormParams {
    users: User[]
    setUsers: React.Dispatch<React.SetStateAction<User[]>>
    open: boolean
    setOpen: React.Dispatch<React.SetStateAction<boolean>>
}

export default function SearchStudentForm({users, setUsers, open, setOpen}: SearchStudentFormParams) {
    const [studentID, setStudentID] = useState('')
    const [firstname, setFirstname] = useState('')
    const [lastname, setLastname] = useState('')

    const closeForm = () => {
        setOpen(false)
    }

    const searchStudent = () => {
        console.log("Search for student...")

        const filters = new Map<() => boolean, () => void>([
            [() => studentID !== "", () => { users = users.filter((user) => user.student_id === studentID)} ],
            [() => firstname !== "", () => { users = users.filter((user) => user.firstname === firstname)} ],
            [() => lastname !== "", () => { users = users.filter((user) => user.lastname === lastname)} ],
        ])

        filters.forEach((filter, check) => {
            if (check()) {
                filter()
            }
        })

        setUsers(users)
    }

    return (
        <Dialog open={open} onClose={closeForm}>
            <DialogTitle>Search for student</DialogTitle>
            <DialogContent>
                <DialogContentText>
                    <TextField autoFocus={true} margin={"dense"} id={"firstname"}
                               label={"Firstname"} type={"text"} fullWidth
                               variant={"outlined"} value={firstname}
                               onChange={(e) => setFirstname(e.target.value)}
                    />

                    <TextField autoFocus={true} margin={"dense"} id={"lastname"}
                               label={"Lastname"} type={"text"} fullWidth
                               variant={"outlined"} value={lastname}
                               onChange={(e) => setLastname(e.target.value)}
                    />

                    <TextField autoFocus={true} margin={"dense"} id={"student_id"}
                               label={"Student ID"} type={"text"} fullWidth variant={"outlined"}
                               value={studentID} onChange={(e) => setStudentID(e.target.value)}
                    />
                </DialogContentText>
            </DialogContent>
            <DialogActions>
                <Button onClick={closeForm} color={"secondary"}>Cancel</Button>
                <Button type="submit" color={"secondary"} onClick={searchStudent}>Search</Button>
            </DialogActions>
        </Dialog>
    )
}