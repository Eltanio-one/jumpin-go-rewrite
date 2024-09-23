import './App.css';
import LoginModal from './components/LoginModal';
import RegisterModal from './components/RegisterModal';
import { GoogleReCaptchaProvider } from "react-google-recaptcha-v3";
import { Box } from '@mui/material';
import GymLoginModal from './components/GymLoginModal';
import GymRegisterModal from './components/GymRegisterModal';

export default function Login() {
    return (
        <div className="App">
            <header className="App-header">
                <p>
                    Welcome to JumpIn
                </p>
                <Box>
                    <GoogleReCaptchaProvider reCaptchaKey="6Lc19UoqAAAAAM2mCg193zC3JghzJDNowoepquia">
                        <LoginModal />
                        <RegisterModal />
                    </GoogleReCaptchaProvider>
                </Box>
                <p style={{ fontSize: '15px' }}>
                    Are you a Gym Owner? Log in below
                </p>
                <Box>
                    <GoogleReCaptchaProvider reCaptchaKey="6Lc19UoqAAAAAM2mCg193zC3JghzJDNowoepquia">
                        <GymLoginModal />
                        <GymRegisterModal />
                    </GoogleReCaptchaProvider>
                </Box>
            </header>
        </div >
    );
}
