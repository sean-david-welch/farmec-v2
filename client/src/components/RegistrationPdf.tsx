import utils from '../styles/Utils.module.css';

import { faDownload } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { MachineRegistration } from '../types/miscTypes';
import { Page, Text, Document, StyleSheet, PDFDownloadLink } from '@react-pdf/renderer';

const styles = StyleSheet.create({
    title: {
        fontSize: 24,
        textAlign: 'center',
        margin: 50,
    },
    content: {
        fontSize: 12,
        textAlign: 'center',
        marginBottom: 5,
        marginHorizontal: 100,
        width: 400,
    },
});

interface Props {
    registration: MachineRegistration;
}

export const DownloadLink = ({ registration }: Props) => (
    <PDFDownloadLink
        document={<MyDocument registration={registration} />}
        fileName={`${registration.owner_name}.registration.pdf`}
    >
        {() => (
            <button className={utils.btn}>
                Download Form
                <FontAwesomeIcon icon={faDownload} />
            </button>
        )}
    </PDFDownloadLink>
);

export const MyDocument = ({ registration }: Props) => (
    <Document>
        <Page size="A4">
            <Text style={styles.title}>
                Machine Registration: {registration.dealer_name} - {registration.owner_name}
            </Text>
            {Object.entries(registration).map(([key, value]) => {
                if (key !== 'id' && key !== 'created' && key !== 'parts') {
                    return (
                        <Text style={styles.content} key={key}>
                            {key}: {String(value)}
                        </Text>
                    );
                }
            })}
        </Page>
    </Document>
);
