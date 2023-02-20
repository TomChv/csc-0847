import React from 'react';
import {createTheme, ThemeProvider} from '@mui/material/styles';
import TopBar from "./components/TopBar/TopBar";
import Footer from "./components/Footer/Footer";
import HomePage from "./components/HomePage/HomePage";

const theme = createTheme({
    palette: {
        primary: {
            main: '#272626'
        },
        secondary: {
            main: '#35EDF9'
        }
    }
})

function App() {
    return (
        <ThemeProvider theme={theme}>
            <TopBar/>
            <HomePage/>
            <Footer/>
        </ThemeProvider>
    );
}

export default App;
