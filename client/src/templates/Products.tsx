import utils from '../styles/Utils.module.css';
import styles from '../styles/Suppliers.module.css';

import { Machine, Product } from '../types/supplierTypes';
import ProductForm from '../forms/ProductForm';
import DeleteButton from '../components/DeleteButton';
import { useGetResource } from '../hooks/genericHooks';

interface Props {
    products: Product[];
    isAdmin: boolean;
}

const Products: React.FC<Props> = ({ products, isAdmin }: Props) => {
    const machines = useGetResource<Machine[]>('machines');

    if (!machines.data) {
        return <div>Loading...</div>;
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

                        {isAdmin && product.id && (
                            <div className={utils.optionsBtn}>
                                <ProductForm id={product.id} product={product} machines={machines.data} />
                                <DeleteButton id={product.id} resourceKey="products" />
                            </div>
                        )}
                    </div>
                ))}
            </div>
            {isAdmin && <ProductForm machines={machines.data} />}
        </section>
    );
};

export default Products;
