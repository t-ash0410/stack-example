/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  images: {
    unoptimized: true,
  },
  publicRuntimeConfig: {
    API_BASE_URL: process.env.API_BASE_URL || 'http://localhost:8080',
    SLACK_CLIENT_ID: process.env.SLACK_CLIENT_ID || '',
  },
}

export default nextConfig
