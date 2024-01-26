import utils from '../../styles/Info.module.css';

const Info: React.FC = () => {
    return (
        <div className={utils.info}>
            <div className={utils.infoSection}>
                <h1 className={utils.subHeading}>Business Information:</h1>
                <div className={utils.infoItem}>
                    Opening Hours:
                    <br />
                    <span className={utils.infoItemText}>Monday - Friday: 9am - 5:30pm</span>
                </div>
                <div className={utils.infoItem}>
                    Telephone:
                    <br />
                    <span className={utils.infoItemText}>
                        <Link to="tel:01 825 9289">01 825 9289</Link>
                    </span>
                </div>
                <div className={utils.infoItem}>
                    International:
                    <br />
                    <span className={utils.infoItemText}>
                        <Link to="tel:+353 1 825 9289">+353 1 825 9289</Link>
                    </span>
                </div>
                <div className={utils.infoItem}>
                    Email:
                    <br />
                    <span className={utils.infoItemText}>Info@farmec.ie</span>
                </div>
                <div className={utils.infoItem}>
                    Address:
                    <br />
                    <span className={utils.infoItemText}>Clonross, Drumree, Co. Meath, A85PK30</span>
                </div>
            </div>
        </div>
    );
};

export default Info;
