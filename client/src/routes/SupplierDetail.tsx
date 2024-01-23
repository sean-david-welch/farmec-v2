import { useParams } from '@solidjs/router';
import { Component, createResource } from 'solid-js';

const SuppliersDetails: Component = () => {
    const params = useParams();
    const [supplier] = createResource(() => {
        return fetch(`http://localhost:8080/api/suppliers/${params.id}`).then(response => response.json());
    });

    return (
        <>
            <h1>Supplier Details</h1>
            {supplier.loading && <p>Loading...</p>}
            {supplier.error && <p>Error loading supplier!</p>}
            {supplier() && (
                // Display your supplier details here
                <div>
                    <h2>{supplier().name}</h2>
                    {/* other details */}
                </div>
            )}
        </>
    );
};

export default SuppliersDetails;
