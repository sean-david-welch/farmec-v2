import utils from '../styles/Utils.module.css';
import styles from '../styles/Account.module.css';

import { LineItem } from '../types/miscTypes';
import { useGetResource } from '../hooks/genericHooks';
import { Fragment } from 'react';
import LineItemForm from '../forms/LineItemForm';
import DeleteButton from '../components/DeleteButton';

const LineItems: React.FC = () => {
    const { data: lineItems } = useGetResource<LineItem[]>('lineitems');

    return (
        <section id="lineItems">
            <Fragment>
                <h1 className={utils.sectionHeading}>Product Line Items:</h1>
                {lineItems &&
                    lineItems.map((lineItem) => (
                        <div className={styles.productView} key={lineItem.id}>
                            <h1 className={utils.mainHeading}>
                                {lineItem.name} -- {lineItem.price}
                            </h1>
                            <img src={lineItem.image} alt={'line item image'} width={400} height={400} />
                            {lineItem.id && (
                                <div className={utils.optionsBtn}>
                                    <LineItemForm id={lineItem.id} lineItem={lineItem} />
                                    <DeleteButton id={lineItem.id} resourceKey="lineitems" />
                                </div>
                            )}
                        </div>
                    ))}
                <LineItemForm />
            </Fragment>
        </section>
    );
};

export default LineItems;
