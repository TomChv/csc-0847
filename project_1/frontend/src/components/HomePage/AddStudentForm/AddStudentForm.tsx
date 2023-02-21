import React, {useState} from "react";
import {
    Button,
    Dialog,
    DialogActions,
    DialogContent,
    DialogContentText,
    DialogTitle,
    TextField
} from "@mui/material";
import {client} from "../../../backend/client";

interface AddStudentFormParam {
    open: boolean
    setOpen: React.Dispatch<React.SetStateAction<boolean>>
}

export default function AddStudentForm({open, setOpen}: AddStudentFormParam) {
    const [firstname, setFirstname] = useState('')
    const [lastname, setLastname] = useState('')
    const [studentID, setStudentID] = useState('')
    const [email, setEmail] = useState('')
    const [mailingAddress, setMailingAddress] = useState('')
    const [gpa, setGPA] = useState('')

    const closeForm = () => {
        setOpen(false)
    }

    const addStudent = () => {
        console.log("AddStudent...")

        console.log(firstname, lastname, studentID, email, mailingAddress, gpa)

        client.createUser({
            firstname,
            lastname,
            student_id: studentID,
            email,
            mailing_address: mailingAddress,
            gpa: Number(gpa),
        })
            .then(() => setOpen(false))
            .catch((e) => console.error(`could not create users ${e.message}`))
    }

    return (
        <Dialog open={open} onClose={closeForm}>
            <DialogTitle>Add Student</DialogTitle>
            <DialogContent>
                <DialogContentText>
                    <TextField autoFocus={true} margin={"dense"} id={"firstname"}
                               label={"Firstname"} type={"text"} fullWidth
                               variant={"outlined"} required value={firstname}
                               onChange={(e) => setFirstname(e.target.value)}
                    />

                    <TextField autoFocus={true} margin={"dense"} id={"lastname"}
                               label={"Lastname"} type={"text"} fullWidth
                               variant={"outlined"} required value={lastname}
                               onChange={(e) => setLastname(e.target.value)}
                    />

                    <TextField autoFocus={true} margin={"dense"} id={"student_id"}
                               label={"Student ID"} type={"text"} fullWidth variant={"outlined"}
                               required value={studentID}
                               onChange={(e) => setStudentID(e.target.value)}
                    />

                    <TextField autoFocus={true} margin={"dense"} id={"email"} label={"Email"}
                               type={"email"} fullWidth variant={"outlined"} required
                               value={email} onChange={(e) => setEmail(e.target.value)}
                    />


                    <TextField autoFocus={true} margin={"dense"} id={"mailing_address"} label={"Mailing Address"}
                               type={"text"} fullWidth variant={"outlined"} required
                               value={mailingAddress}
                               onChange={(e) => setMailingAddress(e.target.value)}
                    />

                    <TextField id="gpa" label="GPA" type="number" InputLabelProps={{shrink: true}}
                               InputProps={{inputProps: {min: 0, max: 4}}} margin={"dense"}
                               variant={"outlined"} required value={gpa}
                               onChange={(e) => setGPA(e.target.value)}
                    />
                </DialogContentText>
            </DialogContent>
            <DialogActions>
                <Button onClick={closeForm} color={"secondary"}>Cancel</Button>
                <Button type="submit" onClick={addStudent} color={"secondary"}>Create</Button>
            </DialogActions>
        </Dialog>
    )
}