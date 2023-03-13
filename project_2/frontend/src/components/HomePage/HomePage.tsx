import React, {useEffect, useState} from "react";
import {Box, Grid, IconButton, Paper} from "@mui/material";
import {Picture} from "../../types/pictures";
import client from "../../backend/backend";
import {picturesMocks} from "../../mocks/pictures";
import AddIcon from '@mui/icons-material/Add';
import PetCard from "./PetCard/PetCard";
import EditPictureForm from "./EditPicture/EditPicture";
import AddPictureForm from "./AddPicture/AddPicture";

export default function HomePage() {
    const [pictures, setPicturesState] = useState<Picture[]>([])
    const [displayAddPictureForm, setDisplayAddPictureForm] = useState(false)
    const [displayEditPictureForm, setDisplayEditPictureForm] = useState(false)
    const [selectedPicToEdit, setSelectedPicToEdit] = useState<Picture | undefined>()
    const [deletePic, setDeletePic] = useState(true)

    useEffect(() => {
        console.log("Fetch pictures")

        client.listPictures()
            .then((pics) => {
                setPicturesState(pics)
                setDeletePic(false)
            })
            .catch((e) => {
                console.log(`could not fetch pictures ${e}`)
                console.log("Use mockup instead")
                setPicturesState(picturesMocks)
                setDeletePic(false)
            })
    }, [deletePic, displayEditPictureForm, displayAddPictureForm])

    return (
        <Box display={"flex"} justifyContent={"center"} alignContent={"center"} margin={"6vh"}>
            <Grid container spacing={{xs: 2, md: 3}} columns={{xs: 4, sm: 8, md: 12}}>
                {pictures.map((picture, index) => (
                    <Grid xs={2} sm={4} md={4} key={index}>
                        <PetCard picture={picture}
                                 setEdit={setDisplayEditPictureForm}
                                 setPicToEdit={setSelectedPicToEdit}
                                 setDelete={setDeletePic}
                        />
                    </Grid>
                ))}
                <Grid xs={2} sm={4} md={4} key={pictures.length}>
                    <Paper  sx={{
                        display: "flex",
                        textAlign: "center",
                        alignItems: "center",
                        justifyContent: "center",
                        height: "300px",
                        margin: "5px",
                        padding: "5px",
                        backgroundColor: "rgba(39,38,38,0.15)"
                    }}>
                        <IconButton onClick={() => setDisplayAddPictureForm(true)}>
                            <AddIcon fontSize={"large"} color={"disabled"}/>
                        </IconButton>
                    </Paper>
                </Grid>
            </Grid>

            <EditPictureForm picture={selectedPicToEdit} open={displayEditPictureForm} setOpen={setDisplayEditPictureForm}/>
            <AddPictureForm open={displayAddPictureForm} setOpen={setDisplayAddPictureForm}/>
        </Box>
    )
}