import utils from '../../styles/Utils.module.css';

const Heading: React.FC = () => {
    return (
        <div className={utils.typewriter}>
            <h1>Importers & Distributors of Quality Agricultural Machinery</h1>

            <button className={utils.btn}>
                <a href="#Info">
                    Find Out More: <img src="/icons/chevron-down.svg" alt="down" />
                </a>
            </button>
        </div>
    );
};

export default Heading;
