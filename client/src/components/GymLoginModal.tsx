import { Alert, Box, Button, Dialog, DialogTitle, TextField } from '@mui/material'
import DialogContent from '@mui/material/DialogContent'
import { useState } from 'react'
import { useGoogleReCaptcha } from "react-google-recaptcha-v3";
import { useNavigate } from 'react-router-dom';
import { useUser } from './UserContext';

export default function GymLoginModal() {

    const { executeRecaptcha } = useGoogleReCaptcha();

    const { setUser } = useUser();

    const [state, setState] = useState<{
        open: boolean,
        usermail: string,
        password: string,
        error: string,
    }>({
        open: false,
        usermail: "",
        password: "",
        error: "",
    })

    const navigate = useNavigate();

    function onOpen() {
        setState((prevState) => ({ ...prevState, open: true }));
    }

    function onClose() {
        setState({
            open: false,
            usermail: "",
            password: "",
            error: "",
        })
    }

    function onTyping(e: React.ChangeEvent<any>) {
        setState((prevState) => ({
            ...prevState,
            [e.target.id]: e.target.value
        }))
    }


    return (
        <>
            <Box>
                <Button variant="outlined" onClick={onOpen} style={{ fontSize: '10px', padding: '5px 10px', borderRadius: '5px' }}>Login</Button>

                <Dialog onClose={onClose} open={state.open} maxWidth="sm" fullWidth slotProps={{
                    backdrop: {
                        style: {
                            backgroundColor: 'rgba(0, 0, 0, 0)',
                        },
                    }
                }}>
                    <DialogTitle>Gym Login</DialogTitle>
                    <DialogContent dividers>
                        <Box>
                            <Box pb={2}>
                                <TextField
                                    fullWidth
                                    label="Username/Email"
                                    variant="outlined"
                                    id="usermail"
                                    value={state.usermail}
                                    onChange={onTyping}
                                />
                            </Box>
                            <Box pb={2}>
                                <TextField
                                    fullWidth
                                    label="Password"
                                    variant="outlined"
                                    id="password"
                                    type="password"
                                    value={state.password}
                                    onChange={onTyping}
                                />
                            </Box>
                        </Box>
                        <Box display="flex" justifyContent="flex-end" alignItems="left" >
                            <Box pr={1}>
                                <Button variant="outlined" onClick={onClose}>Cancel</Button>
                            </Box>
                            <Box pr={1}>
                                <Button variant="outlined">Login</Button>
                            </Box>
                        </Box>
                        <Box></Box>
                        {!!state.error && (
                            <Box mt={2}>
                                <Alert severity="error" variant="outlined">{state.error}</Alert>
                            </Box>
                        )}
                    </DialogContent>
                </Dialog>
            </Box>
        </>
    )
}