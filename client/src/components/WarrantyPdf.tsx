import utils from '../styles/Utils.module.css';

import { PartsRequired, WarrantyClaim } from '../types/miscTypes';
import { Page, Text, Document, StyleSheet } from '@react-pdf/renderer';

import { PDFDownloadLink } from '@react-pdf/renderer';

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
    warranty: WarrantyClaim;
    parts: PartsRequired[];
}

export const DownloadLink = ({ warranty, parts }: Props) => (
    <PDFDownloadLink
        document={<MyDocument warranty={warranty} parts={parts} />}
        fileName={`${warranty.owner_name}.warranty.pdf`}>
        {() => (
            <button className={utils.btn}>
                Download Form
                <img src="/icons/download.svg" alt="download" />
            </button>
        )}
    </PDFDownloadLink>
);

export const MyDocument = ({ warranty, parts }: Props) => (
    <Document>
        <Page size="A4">
            <Text style={styles.title}>
                Warranty Claim: {warranty.dealer} - {warranty.owner_name}
            </Text>
            {Object.entries(warranty).map(([key, value]) => {
                if (key !== 'id' && key !== 'created' && key !== 'parts') {
                    return (
                        <Text style={styles.content} key={key}>
                            {key}: {String(value)}
                        </Text>
                    );
                }
            })}
            {parts.map((part, index) =>
                Object.entries(part).map(([key, value]) => {
                    if (key !== 'id' && key !== 'warrantyId') {
                        return (
                            <Text style={styles.content} key={key + index}>
                                {key}: {String(value)}
                            </Text>
                        );
                    }
                })
            )}
        </Page>
    </Document>
);
