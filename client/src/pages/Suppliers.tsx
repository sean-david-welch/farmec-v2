import styles from '../styles/Suppliers.module.css';
import utils from '../styles/Utils.module.css';

import SupplierForm from '../forms/SupplierForm';
import DeleteButton from '../components/DeleteButton';

import {Link} from 'react-router-dom';
import {Fragment} from 'react';

import {useUserStore} from '../lib/store';
import {SocialLinks} from '../components/SocialLinks';
import {useGetResource} from '../hooks/genericHooks';
import {Supplier} from '../types/supplierTypes';
import ErrorPage from '../layouts/Error';
import Loading from '../layouts/Loading';
import {Helmet} from "react-helmet";

const Suppliers: React.FC = () => {
    const {isAdmin} = useUserStore();
    const {data: suppliers, isError, isLoading} = useGetResource<Supplier[]>('suppliers');

    if (isError) return <ErrorPage/>;
    if (isLoading) return <Loading/>;

    const imageError = (event: React.SyntheticEvent<HTMLImageElement, Event>) => {
        event.currentTarget.src = '/default.jpg';
    };

    return (
        <>
            <Helmet>
                <title>Suppliers - Farmec Ireland</title>
                <meta name="description" content="Browse our Suppliers and learn more about the machines we offer."/>

                <meta property="og:title" content="Suppliers - Farmec Ireland"/>
                <meta property="og:description" content="Browse our Suppliers and learn more about the machines we offer."/>
                <meta property="og:image" content="https://www.farmec.ie/farmec_images/Suppliers/sip1250.webp"/>
                <meta property="og:url" content="https://www.farmec.ie/suppliers"/>
                <meta property="og:type" content="website"/>

                <meta name="twitter:card" content="summary_large_image"/>
                <meta name="twitter:title" content="Suppliers - Farmec Ireland"/>
                <meta name="twitter:description" content="Browse our Suppliers and learn more about the machines we offer."/>
                <meta name="twitter:image" content="https://www.farmec.ie/farmec_images/Suppliers/sip1250.webp"/>
                <link rel="canonical" href="https://www.farmec.ie/suppliers"/>
            </Helmet>
            <section id="suppliers">
                <h1 className={utils.sectionHeading}>Suppliers</h1>

                {isAdmin && <SupplierForm/>}
                {suppliers ? (
                    <div className={utils.index}>
                        <h1 className={utils.indexHeading}>Suppliers:</h1>
                        {suppliers.map(link => (
                            <a key={link.name} href={`#${link.name}`}>
                                <h1 className={utils.indexItem}>{link.name}</h1>
                            </a>
                        ))}
                    </div>
                ) : null}

                {Array.isArray(suppliers) &&
                    suppliers.map(supplier => (
                        <Fragment key={supplier.id}>
                            <div className={styles.supplierCard} id={supplier.name}>
                                <div className={styles.supplierGrid}>
                                    <div className={styles.supplierHead}>
                                        <h1 className={utils.mainHeading}>{supplier.name}</h1>
                                        <img
                                            src={supplier.logo_image}
                                            alt="Supplier logo"
                                            className={styles.supplierLogo}
                                            width={200}
                                            height={200}
                                            onError={imageError}
                                        />

                                        <SocialLinks
                                            facebook={supplier.social_facebook}
                                            twitter={supplier.social_twitter}
                                            instagram={supplier.social_instagram}
                                            linkedin={supplier.social_linkedin}
                                            website={supplier.social_website}
                                            youtube={supplier.social_youtube}
                                        />
                                    </div>
                                    <img
                                        src={supplier.marketing_image}
                                        alt="Marketing"
                                        className={styles.supplierImage}
                                        width={550}
                                        height={550}
                                        onError={imageError}
                                    />
                                </div>
                                <div className={styles.supplierInfo}>
                                    <p className={styles.supplierDescription}>{supplier.description}</p>
                                    <button className={utils.btn}>
                                        <Link to={`/suppliers/${supplier.id}`}>Learn More</Link>
                                    </button>
                                </div>
                            </div>

                            {isAdmin && supplier.id && (
                                <div className={utils.optionsBtn}>
                                    <SupplierForm id={supplier.id} supplier={supplier}/>
                                    <DeleteButton id={supplier.id} resourceKey="suppliers"/>
                                </div>
                            )}
                        </Fragment>
                    ))}
            </section>
        </>
    );
};

export default Suppliers;
