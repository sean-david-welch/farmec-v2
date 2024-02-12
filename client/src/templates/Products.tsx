import utils from '../styles/Utils.module.css';
import styles from '../styles/Suppliers.module.css';

import Loading from '../layouts/Loading';
import ProductForm from '../forms/ProductForm';
import DeleteButton from '../components/DeleteButton';

import { Machine, Product } from '../types/supplierTypes';
import { useGetResourceById } from '../hooks/genericHooks';

interface Props {
    id: string;
    isAdmin: boolean;
    products: Product[];
}

const Products: React.FC<Props> = ({ id, isAdmin, products }: Props) => {
    const { data: machine } = useGetResourceById<Machine>('machines', id);

    if (!machine) {
        return <Loading />;
    }

    return (
        <section id="products">
            <div className={styles.productGrid}>
                {products.map(product => (
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
                                <ProductForm id={product.id} product={product} machine={machine} />
                                <DeleteButton id={product.id} resourceKey="products" />
                            </div>
                        )}
                    </div>
                ))}
            </div>
            {isAdmin && <ProductForm machine={machine} />}
        </section>
    );
};

export default Products;
