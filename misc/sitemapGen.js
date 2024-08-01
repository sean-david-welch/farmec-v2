
import axios from 'axios';
import fs from 'fs';
import { create } from 'xmlbuilder';

const BASE_URL = 'https://www.farmec.ie';
const OUTPUT_FILE = 'sitemap.xml';

// Static routes
const staticRoutes = [
  '/',
  '/about',
  '/about/policies',
  '/suppliers',
  '/spareparts',
  '/blogs',
  '/blog/exhibitions',
  '/warranties',
  '/registrations',
  '/login',
  '/return',
];

// Example function to fetch dynamic routes
async function fetchDynamicRoutes() {
  // Mock data fetching functions. Replace with actual API calls or data fetching logic.
  const supplierIds = await axios.get(`${BASE_URL}/api/suppliers`); // Replace with actual endpoint
  const sparePartIds = await axios.get(`${BASE_URL}/api/spareparts`); // Replace with actual endpoint

  // Return dynamic routes as an array of paths
  return {
    suppliers: supplierIds.data.map(id => `/suppliers/${id}`),
    spareparts: sparePartIds.data.map(id => `/spareparts/${id}`),
  };
}

// Function to generate sitemap
async function generateSitemap() {
  try {
    const dynamicRoutes = await fetchDynamicRoutes();

    // Create XML structure
    const urlset = create('urlset', {
      version: '1.0',
      encoding: 'UTF-8',
    }).att('xmlns', 'http://www.sitemaps.org/schemas/sitemap/0.9');

    // Add static routes
    staticRoutes.forEach(route => {
      urlset.ele('url')
        .ele('loc', `${BASE_URL}${route}`).up()
        .ele('lastmod', new Date().toISOString());
    });

    // Add dynamic routes for suppliers
    dynamicRoutes.suppliers.forEach(route => {
      urlset.ele('url')
        .ele('loc', `${BASE_URL}${route}`).up()
        .ele('lastmod', new Date().toISOString());
    });

    // Add dynamic routes for spare parts
    dynamicRoutes.spareparts.forEach(route => {
      urlset.ele('url')
        .ele('loc', `${BASE_URL}${route}`).up()
        .ele('lastmod', new Date().toISOString());
    });

    // Generate XML string and write to file
    const xml = urlset.end({ pretty: true });
    fs.writeFileSync(OUTPUT_FILE, xml, 'utf8');
    console.log(`Sitemap successfully generated and saved to ${OUTPUT_FILE}`);
  } catch (error) {
    console.error('Error generating sitemap:', error);
  }
}

// Run the script
generateSitemap();
