import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import utils from '../styles/Utils.module.css';
import { faDownload } from '@fortawesome/free-solid-svg-icons';
import { MachineRegistration, WarrantyParts } from '../types/miscTypes';
import config from '../lib/env';

interface Props {
    warrantyClaim?: WarrantyParts;
    registration?: MachineRegistration;
}

const DownloadPdfButton: React.FC<Props> = ({ warrantyClaim, registration }) => {
    async function handleSubmit(event: React.MouseEvent<HTMLButtonElement>) {
        event.preventDefault();

        let endpoint = '';
        let body = {};

        if (warrantyClaim) {
            endpoint = 'warranty';
            const { warranty, parts } = warrantyClaim;

            const partsRequired = parts.map(part => {
                return {
                    part_number: part.part_number,
                    quantity_needed: part.quantity_needed,
                    invoice_number: part.invoice_number,
                    description: part.description,
                };
            });

            body = {
                warranty: {
                    dealer: warranty.dealer,
                    dealer_contact: warranty.dealer_contact,
                    owner_name: warranty.owner_name,
                    machine_model: warranty.machine_model,
                    serial_number: warranty.serial_number,
                    install_date: warranty.install_date,
                    failure_date: warranty.failure_date,
                    repair_date: warranty.repair_date,
                    failure_details: warranty.failure_details,
                    repair_details: warranty.repair_details,
                    labour_hours: warranty.labour_hours,
                    completed_by: warranty.completed_by,
                },
                parts: partsRequired,
            };
        } else if (registration) {
            endpoint = 'registration';

            body = {
                dealer_name: registration.dealer_name,
                dealer_address: registration.dealer_address,
                owner_name: registration.owner_name,
                owner_address: registration.owner_address,
                machine_model: registration.machine_model,
                serial_number: registration.serial_number,
                install_date: registration.install_date,
                invoice_number: registration.invoice_number,
                complete_supply: registration.complete_supply,
                pdi_complete: registration.pdi_complete,
                pto_correct: registration.pto_correct,
                machine_test_run: registration.machine_test_run,
                safety_induction: registration.safety_induction,
                operator_handbook: registration.operator_handbook,
                date: registration.date,
                completed_by: registration.completed_by,
            };
        }

        const url = new URL(`api/pdf/${endpoint}`, config.baseUrl);

        const triggerDownload = (blob: Blob, filename: string) => {
            const downloadUrl = window.URL.createObjectURL(blob);
            const link = document.createElement('a');

            link.href = downloadUrl;
            link.download = filename;
            document.body.appendChild(link);

            link.click();

            document.body.removeChild(link);
            window.URL.revokeObjectURL(downloadUrl);
        };

        try {
            const response = await fetch(url.toString(), {
                method: 'POST',
                credentials: 'include',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(body),
            });
            if (!response.ok) throw new Error('Network response was not ok');

            const blob = await response.blob();
            triggerDownload(blob, `${endpoint}.pdf`);
        } catch (error) {
            console.error('Error downloading the file:', error);
        }
    }

    return (
        <button className={utils.btn} onClick={handleSubmit}>
            Download Form
            <FontAwesomeIcon icon={faDownload} />
        </button>
    );
};

export default DownloadPdfButton;
