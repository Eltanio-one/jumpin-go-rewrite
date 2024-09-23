import JumpInModal from "./components/JumpInModal";
import { useUser } from "./components/UserContext";
import { Box } from '@mui/material';

export default function HomePage() {

    const { username } = useUser();

    return (
        <div className="App">
            <header className="App-header">
                <p>
                    Hello, {username}
                </p>
                <Box>
                    <JumpInModal />
                </Box>
            </header>
        </div>
    )
}