import styles from '../styles/Utils.module.css';

import LoginForm from '../forms/LoginForm';

const Login: React.FC = () => {
    return (
        <div className={styles.loginSection}>
            <LoginForm />
        </div>
    );
};

export default Login;
