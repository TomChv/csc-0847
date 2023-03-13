import React, {useState} from "react";
import {
    Button,
    Dialog,
    DialogActions,
    DialogContent,
    DialogContentText,
    DialogTitle,
    IconButton,
    TextField, Tooltip
} from "@mui/material";
import {PhotoCamera} from "@mui/icons-material";
import client from "../../../backend/backend";

interface AddPictureFormParam {
    open: boolean
    setOpen: React.Dispatch<React.SetStateAction<boolean>>
}

export default function AddPictureForm({ open, setOpen }: AddPictureFormParam) {
    const [date, setDate] = useState('')
    const [location, setLocation] = useState('')
    const [picture, setPicture] = useState<File | null>(null)
    const [author, setAuthor] = useState('')

    const handleCapture = ({ target }: any) => {
        setPicture(target.files[0]);
    };

    const addPicture = () => {
        console.log("Add picture...")

        console.log(picture, date, location, author)
        if (!picture) {
            console.log("No file inserted")
            return
        }

        client.addPicture(picture, { date, location, author })
            .then(() => console.log("picture updated"))
            .catch((e) => console.log(`could not update picture ${e}`))
            .finally(() => setOpen(false))
    }

    return (
        <Dialog open={open} onClose={() => setOpen(false)}>
            <DialogTitle>Add picture</DialogTitle>
            <DialogContent>
                <DialogContentText>
                    <TextField autoFocus={true} margin={"dense"} id={"date"}
                               label={"Date"} type={"date"} fullWidth
                               variant={"outlined"} value={date} required
                               onChange={(e) => setDate(e.target.value)}
                    />

                    <TextField autoFocus={true} margin={"dense"} id={"location"}
                               label={"Location"} type={"text"} fullWidth
                               variant={"outlined"} value={location} required
                               onChange={(e) => setLocation(e.target.value)}
                    />

                    <TextField autoFocus={true} margin={"dense"} id={"author"}
                               label={"Author"} type={"text"} fullWidth variant={"outlined"}
                               value={author} required
                               onChange={(e) => setAuthor(e.target.value)}
                    />

                    <input
                        accept="image/jpeg"
                        id="faceImage"
                        type="file"
                        hidden={true}
                        onChange={handleCapture}
                    />
                    <Tooltip title="Select Image">
                        <label htmlFor="faceImage">
                            <IconButton
                                color="primary"
                                aria-label="upload picture"
                                component="span"
                            >
                                <PhotoCamera fontSize="large" />
                            </IconButton>
                        </label>
                    </Tooltip>
                    <label>{picture ? picture.name : "Select Image"}</label>

                </DialogContentText>
            </DialogContent>
            <DialogActions>
                <Button onClick={() => setOpen(false)} color={"secondary"}>Cancel</Button>
                <Button type="submit" onClick={addPicture} color={"secondary"}>Add</Button>
            </DialogActions>
        </Dialog>
    )
}