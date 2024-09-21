import { Box, Button, Dialog, DialogTitle, TextField } from '@mui/material'
import DialogContent from '@mui/material/DialogContent'
import { useState } from 'react'

export default function LoginModal() {

    const [state, setState] = useState<{
        open: boolean,
        usermail: string,
        password: string,
    }>({
        open: false,
        usermail: "",
        password: "",
    })

    function onOpen() {
        setState((prevState) => ({ ...prevState, open: true }));
    }

    function onClose() {
        setState({
            open: false,
            usermail: "",
            password: "",
        })
    }

    function onTyping(e: React.ChangeEvent<any>) {
        setState((prevState) => ({
            ...prevState,
            [e.target.id]: e.target.value
        }))
    }

    function onSubmit() {
        fetch("http://localhost:8080/login", {
            headers: {
                Accept: "application/json",
                "Content-Type": "application/json",
            },
            method: "POST",
            body: JSON.stringify({
                usermail: state.usermail,
                password: state.password,
            }),
        })
            .then(async (resp) => {
                if (resp.status === 200) {
                    onClose()
                    return
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
                        <Box display="flex" justifyContent="flex-end">
                            <Box pr={1}>
                                <Button variant="outlined" onClick={onClose}>Cancel</Button>
                            </Box>
                            <Box pr={1}>
                                <Button variant="outlined" onClick={onSubmit}>Login</Button>
                            </Box>
                        </Box>
                    </DialogContent>
                </Dialog>
            </Box>
        </>
    )

}
