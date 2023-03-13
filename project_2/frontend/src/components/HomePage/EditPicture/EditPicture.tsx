import { Picture } from "../../../types/pictures";
import React, {useState} from "react";
import {
    Button,
    Dialog,
    DialogActions,
    DialogContent,
    DialogContentText,
    DialogTitle, InputLabel,
    MenuItem, Select,
    TextField
} from "@mui/material";
import client from "../../../backend/backend";

interface EditPictureFormParam {
    picture: Picture | undefined,
    open: boolean
    setOpen: React.Dispatch<React.SetStateAction<boolean>>
}

export default function EditPictureForm({ picture, open, setOpen }: EditPictureFormParam) {
    const [date, setDate] = useState(picture?.date)
    const [location, setLocation] = useState(picture?.location)
    const [label, setLabel] = useState<string[]>([picture?.label || ""])
    const [author, setAuthor] = useState(picture?.author)

    const editPicture = () => {
        console.log("Edit picture...")

        console.log(picture?.name, date, location, label, author)

        if (!picture?.name) {
            console.log("no picture selected")
            return
        }

        client.editPicture(picture.name, { date, location, label: label[0], author })
            .then(() => {
                console.log("picture updated")
            })
            .catch((e) => {
                console.log(`could not update picture ${e}`)
            })
            .finally(() => setOpen(false))

    }

    return (
        <Dialog open={open} onClose={() => setOpen(false)}>
            <DialogTitle>Edit {picture?.name}</DialogTitle>
            <DialogContent>
                <DialogContentText>
                    <TextField autoFocus={true} margin={"dense"} id={"date"}
                               label={"Date"} type={"date"} fullWidth
                               variant={"outlined"} value={date}
                               onChange={(e) => setDate(e.target.value)}
                    />

                    <TextField autoFocus={true} margin={"dense"} id={"location"}
                               label={"Location"} type={"text"} fullWidth
                               variant={"outlined"} value={location}
                               onChange={(e) => setLocation(e.target.value)}
                    />

                    <TextField autoFocus={true} margin={"dense"} id={"author"}
                               label={"Author"} type={"text"} fullWidth variant={"outlined"}
                               value={author}
                               onChange={(e) => setAuthor(e.target.value)}
                    />

                    <InputLabel id="label">Label</InputLabel>
                    <Select
                        id="label"
                        value={label[0]}
                        margin={"dense"}
                        autoFocus={true}
                        fullWidth
                        variant={"outlined"}
                        onChange={(e) => setLabel([e.target.value])}
                    >
                        <MenuItem value={"Dog"}>Dog</MenuItem>
                        <MenuItem value={"Flower"}>Flower</MenuItem>
                        <MenuItem value={"Person"}>Person</MenuItem>
                        <MenuItem value={"Other"}>Other</MenuItem>
                    </Select>
                </DialogContentText>
            </DialogContent>
            <DialogActions>
                <Button onClick={() => setOpen(false)} color={"secondary"}>Cancel</Button>
                <Button type="submit" onClick={editPicture} color={"secondary"}>Edit</Button>
            </DialogActions>
        </Dialog>
    )
}