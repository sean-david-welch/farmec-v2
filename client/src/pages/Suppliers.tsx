import styles from '../styles/Suppliers.module.css';
import utils from '../styles/Utils.module.css';

import SupplierForm from '../forms/SupplierForm';
import DeleteButton from '../components/DeleteButton';

import { Link } from 'react-router-dom';
import { Fragment } from 'react';
import { Supplier } from '../types/supplierTypes';
import { useUserStore } from '../lib/store';
import { useGetResource } from '../hooks/genericHooks';
import { SocialLinks } from '../components/SocialLinks';

const Suppliers: React.FC = () => {
    const { isAdmin } = useUserStore();
    const suppliers = useGetResource<Supplier[]>('suppliers');

    if (suppliers.isLoading) {
        return <div>Loading...</div>;
    }

    if (suppliers.isError) {
        return <div>Error loading data</div>;
    }

    return (
        <section id="suppliers">
            {suppliers.data && (
                <div className={utils.index}>
                    <h1 className={utils.indexHeading}>Suppliers</h1>
                    {suppliers.data.map((link) => (
                        <a key={link.name} href={`#${link.name}`}>
                            <h1 className="indexItem">{link.name}</h1>
                        </a>
                    ))}
                </div>
            )}

            {suppliers.data
                ? suppliers.data.map((supplier) => (
                      <Fragment key={supplier.id}>
                          <div className={styles.supplierCard} id={supplier.name}>
                              <div className={styles.supplierGrid}>
                                  <div className={styles.supplierHead}>
                                      <h1 className={styles.mainHeading}>{supplier.name}</h1>
                                      <img
                                          src={supplier.logo_image || '/default.jpg'}
                                          alt={'/default.jpg'}
                                          className={styles.supplierLogo}
                                          width={200}
                                          height={200}
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
                                      src={supplier.marketing_image || '/default.jpg'}
                                      alt={'/default.jpg'}
                                      className={styles.supplierImage}
                                      width={550}
                                      height={550}
                                  />
                              </div>
                              <div className={styles.supplierInfo}>
                                  <p className={styles.supplierDescription}>{supplier.description}</p>
                                  <button className={styles.btn}>
                                      <Link to={`/suppliers/${supplier.id}`}>Learn More</Link>
                                  </button>
                              </div>
                          </div>

                          {isAdmin && supplier.id && (
                              <div className={utils.optionsBtn}>
                                  <SupplierForm id={supplier.id} supplier={supplier} />
                                  <DeleteButton id={supplier.id} resourceKey="suppliers" />
                              </div>
                          )}
                      </Fragment>
                  ))
                : null}

            {isAdmin && <SupplierForm />}
        </section>
    );
};

export default Suppliers;
