import styles from '../styles/Account.module.css';

import LoginForm from '../forms/LoginForm';

const Login: React.FC = () => {
    return (
        <div className={styles.accountMap}>
            <div className={styles.accountSection}>
                <LoginForm />
            </div>
        </div>
    );
};

export default Login;
