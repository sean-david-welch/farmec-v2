import styles from '../styles/Account.module.css';

import LoginForm from '../forms/LoginForm';
import LogoutForm from '../forms/LogoutForm';

const Login: React.FC = () => {
    return (
        <div className={styles.accountMap}>
            <div className={styles.accountSection}>
                <LoginForm />
                <LogoutForm />
            </div>
        </div>
    );
};

export default Login;
