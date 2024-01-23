import { Component } from 'solid-js';
import { createResource, For } from 'solid-js';

const Suppliers: Component = () => {
    const [suppliers] = createResource(async () => {
        const response = await fetch('http://localhost:8080/api/suppliers');
        return response.json();
    });

    return (
        <>
            <h1>Suppliers</h1>
            <ul>
                <For each={suppliers()}>
                    {supplier => (
                        <li>
                            {supplier.name} -- {supplier.id}
                        </li>
                    )}
                </For>
            </ul>
        </>
    );
};

export default Suppliers;
