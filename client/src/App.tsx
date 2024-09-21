import './App.css';
import LoginModal from './components/LoginModal';
import { GoogleReCaptchaProvider } from "react-google-recaptcha-v3";

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <p>
          Welcome to JumpIn
        </p>
        <GoogleReCaptchaProvider reCaptchaKey="6Lc19UoqAAAAAM2mCg193zC3JghzJDNowoepquia">
          <LoginModal />
        </GoogleReCaptchaProvider>
      </header>
    </div>
  );
}

export default App;
