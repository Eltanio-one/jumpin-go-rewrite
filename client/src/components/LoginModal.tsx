import { Alert, Box, Button, Dialog, DialogTitle, TextField } from '@mui/material'
import DialogContent from '@mui/material/DialogContent'
import { useState } from 'react'
import { useGoogleReCaptcha } from "react-google-recaptcha-v3";

export default function LoginModal() {

    const { executeRecaptcha } = useGoogleReCaptcha();

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

    async function onSubmit() {
        if (!executeRecaptcha) {
            setState((prevState) => ({ ...prevState, error: 'recaptcha not yet available' }));
            return;
        }

        const token = await executeRecaptcha('login');

        fetch("http://localhost:8080/login", {
            method: "POST",
            body: JSON.stringify({
                usermail: state.usermail,
                password: state.password,
                recaptcha_token: token,
            }),
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
                <Button variant="outlined" onClick={onOpen}>Login</Button>

                <Dialog onClose={onClose} open={state.open} maxWidth="sm" fullWidth slotProps={{
                    backdrop: {
                        style: {
                            backgroundColor: 'rgba(0, 0, 0, 0)',
                        },
                    }
                }}>
                    <DialogTitle>Login</DialogTitle>
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
                                <Button variant="outlined" onClick={onSubmit}>Login</Button>
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
