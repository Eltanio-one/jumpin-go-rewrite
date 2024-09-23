import './App.css';
import LoginModal from './components/LoginModal';
import RegisterModal from './components/RegisterModal';
import { GoogleReCaptchaProvider } from "react-google-recaptcha-v3";
import { Box } from '@mui/material';

function App() {
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
      </header>
    </div>
  );
}

export default App;
