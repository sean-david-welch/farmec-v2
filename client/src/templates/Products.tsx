import utils from '../styles/Utils.module.css';
import styles from '../styles/Suppliers.module.css';

import { Product } from '../types/supplierTypes';

interface Props {
    products: Product[];
    isAdmin: boolean;
}

const Products: React.FC<Props> = ({ products, isAdmin }: Props) => {
    if (!products) {
        return <div>loading...</div>;
    }

    return (
        <section id="products">
            <div className={styles.productGrid}>
                {products.map((product) => (
                    <div className={styles.productCard} key={product.id} id={product.name || ''}>
                        <h1 className={utils.mainHeading}>{product.name}</h1>
                        <a href={product.product_link || '#'} target="_blank">
                            <img
                                src={product.product_image || '/default.jpg'}
                                alt={'/default.jpg'}
                                className={styles.productImage}
                                width={500}
                                height={500}
                            />
                        </a>
                        <p className={utils.paragraph}>{product.description}</p>
                        {/* {isAdmin && <UpdateProduct product={product} />} */}
                    </div>
                ))}
            </div>
            {/* {isAdmin && <ProductForm machines={} />} */}
        </section>
    );
};

export default Products;