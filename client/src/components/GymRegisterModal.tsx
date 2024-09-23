import { Alert, Box, Button, Dialog, DialogTitle, TextField } from '@mui/material'
import DialogContent from '@mui/material/DialogContent'
import { useState } from 'react'
import DatePicker from "react-datepicker";
import "react-datepicker/dist/react-datepicker.css";

export default function GymRegisterModal() {

    const [state, setState] = useState<{
        open: boolean,
        name: string,
        email: string,
        password: string,
        confirmation: string,
        address: string,
        error: string,
    }>({
        open: false,
        name: "",
        email: "",
        password: "",
        confirmation: "",
        address: "",
        error: "",
    })

    function onOpen() {
        setState((prevState) => ({ ...prevState, open: true }));
    }

    function onClose() {
        setState((prevState) => ({
            ...prevState,
            open: false,
            name: "",
            email: "",
            password: "",
            confirmation: "",
            address: "",
            error: "",
        }))
    }

    function onTyping(e: React.ChangeEvent<any>) {
        setState((prevState) => ({
            ...prevState,
            [e.target.id]: e.target.value
        }))
    }

    async function onSubmit() {

        fetch("http://localhost:8080/gymregister", {
            method: "POST",
            body: JSON.stringify({
                name: state.name,
                email: state.email,
                password: state.password,
                confirmation: state.confirmation,
                address: state.address,
            })
        })
            .then(async (resp) => {
                if (resp.status === 200) {
                    onClose()
                    return
                } else {
                    let errorMessage = await resp.text();
                    setState((prevState) => ({
                        ...prevState,
                        error: errorMessage,
                    })
                    )
                }
            })
            .catch((err) => console.log(err))
    }


    return (
        <>
            <Box>
                <Button variant="outlined" onClick={onOpen} style={{ fontSize: '10px', padding: '5px 10px', borderRadius: '5px' }}>Register</Button>

                <Dialog onClose={onClose} open={state.open} maxWidth="sm" fullWidth slotProps={{
                    backdrop: {
                        style: {
                            backgroundColor: 'rgba(0, 0, 0, 0)',
                        },
                    }
                }}>
                    <DialogTitle>Gym Register</DialogTitle>
                    <DialogContent dividers>
                        <Box>
                            <Box pb={2}>
                                <TextField
                                    fullWidth
                                    label="Gym Name"
                                    variant="outlined"
                                    id="name"
                                    value={state.name}
                                    onChange={onTyping}
                                />
                            </Box>
                            <Box pb={2}>
                                <TextField
                                    fullWidth
                                    label="Gym Email"
                                    variant="outlined"
                                    id="email"
                                    value={state.email}
                                    onChange={onTyping}
                                />
                            </Box>
                            <Box pb={2}>
                                <TextField
                                    fullWidth
                                    label="Address"
                                    variant="outlined"
                                    id="address"
                                    value={state.address}
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
                                /> (Passwords must contain at least 8 characters (one uppercase), one number and one special character)
                            </Box>
                            <Box pb={2}>
                                <TextField
                                    fullWidth
                                    label="Confirm Password"
                                    variant="outlined"
                                    id="confirmation"
                                    type="password"
                                    value={state.confirmation}
                                    onChange={onTyping}
                                />
                            </Box>
                        </Box>
                        <Box display="flex" justifyContent="flex-end" alignItems="left" >
                            <Box pr={1}>
                                <Button variant="outlined" onClick={onClose}>Cancel</Button>
                            </Box>
                            <Box pr={1}>
                                <Button variant="outlined" onClick={onSubmit}>Register</Button>
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
