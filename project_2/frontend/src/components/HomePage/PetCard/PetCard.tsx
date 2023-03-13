import {Picture} from "../../../types/pictures";
import {Box, IconButton, Paper, Typography} from "@mui/material";
import {Image} from "mui-image";
import EditIcon from "@mui/icons-material/Edit";
import DeleteIcon from '@mui/icons-material/Delete';
import React from "react";
import client from "../../../backend/backend";

interface PetCardParameters {
    picture: Picture
    setEdit: React.Dispatch<React.SetStateAction<boolean>>
    setPicToEdit: React.Dispatch<React.SetStateAction<Picture | undefined>>
    setDelete: React.Dispatch<React.SetStateAction<boolean>>
}

export default function PetCard({ picture, setEdit, setPicToEdit, setDelete }: PetCardParameters) {
    const deletePicture = () => {
        client.deletePicture(picture.name)
            .then(() => console.log("picture deleted"))
            .catch((e) => console.log(`could not delete picture ${e}`))

        setDelete(true)
    }

    return (
        <Paper sx={{
            display: "flex",
            flexDirection: "column",
            alignContent: "center",
            justifyContent: "center",
            padding: "5px",
            height: "330px",
            margin: "5px",
        }}>
            <Box>
                <Image height={"200px"} width={"360px"} src={picture.url}/>
                <Typography variant={"body1"}>{picture.name.substring(0, picture.name.length - 4)} by {picture.author} in {picture.location}</Typography>
                <Typography variant={"body1"}></Typography>
                <Typography variant={"body1"}>Posted at {picture.date}</Typography>
                <Typography variant={"body1"}>Label: {picture.label}</Typography>
                <IconButton onClick={() => {
                    setEdit(true)
                    setPicToEdit(picture)
                }}>
                    <EditIcon/>
                </IconButton>
                <IconButton onClick={() => {
                    deletePicture()
                }}>
                    <DeleteIcon/>
                </IconButton>
            </Box>
        </Paper>
    )
}