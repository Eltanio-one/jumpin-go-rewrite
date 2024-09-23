import { useState } from "react";
import { useUser } from "./UserContext";
import { Box, Button, Dialog, DialogContent, DialogTitle, TextField } from "@mui/material";
import { useNavigate } from "react-router-dom";

export default function JumpInModal() {

    const { username } = useUser();

    const [state, setState] = useState<{
        open: boolean,
        username: string | null,
        machines: string[] | null,
    }>({
        open: false,
        username: username,
        machines: null,
    })

    const navigate = useNavigate();

    function onOpen() {
        setState((prevState) => ({ ...prevState, open: true }));
    }

    function onClose() {
        setState((prevState) => ({ ...prevState, open: false }))
    }

    function onTyping() {
        setState((prevState) => ({ ...prevState, }))
    }

    return (
        <>
            <Button variant="outlined" onClick={onOpen}>Jump In</Button>

            <Dialog onClose={onClose} open={state.open} maxWidth="sm" fullWidth slotProps={{
                backdrop: {
                    style: {
                        backgroundColor: 'rgba(0, 0, 0, 0)',
                    },
                }
            }}>
                <DialogTitle>Jump In</DialogTitle>
                <DialogContent dividers>
                    <Box> Please Choose The Machines You Want To Use In The Order You Want To Use Them
                        <Box pb={2}>
                            <TextField
                                fullWidth
                                label="Username/Email"
                                variant="outlined"
                                id="usermail"
                            />
                        </Box>
                    </Box>
                </DialogContent>
            </Dialog>
        </>
    )
}