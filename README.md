# Live Streaming System with RTMP and HLS

## Project Overview

This project is a **live streaming system** that leverages the **RTMP (Real-Time Messaging Protocol)** and **HLS (HTTP Live Streaming)** standards for seamless and real-time video streaming capabilities. The system is designed to provide scalable, high-quality video streaming services using modern and efficient technologies.

The architecture integrates components built in **Golang**, **JavaScript**, **Docker**, and **MySQL**, ensuring a solid backend, a dynamic frontend, and smooth deployment with containerized services.

---

## Key Features

- **RTMP Ingestion**: Accepts live streams via RTMP for encoding and delivery.
- **HLS Output**: Generates HLS streams (.m3u8) for easy consumption by video players.
- **Multi-Resolution Support**: Supports adaptive streaming with configurable video resolutions and bitrates.
- **Redis Integration**: Caches settings and configuration for low-latency, efficient validation.
- **MySQL Database**: Manages stream metadata, user authentication, and other system data.
- **Scalable Deployment**: Containerized using Docker for scalability and portability.
- **Frontend Panel**: A modern web interface to control and monitor live streams.

---

## Tech Stack

### Backend
- **Golang**: Core processing, live encoding, and RTMP-to-HLS conversion.
- **Redis**: For caching system settings and reducing database overhead.
- **MySQL**: Persists application and stream metadata.
- **FFmpeg**: Used for processing and encoding streams dynamically.
- **JS/Socket.IO**: For chat system

### Frontend
- **React + TS**
- **Shacdn**

### Infrastructure
- **Docker**: Containerizes the application for simplicity, scalability, and quick deployments.
- **GitHub Action**: For CI/CD



---

## Architecture Overview

1. **RTMP Ingestion**:
    - Streamers push their live content to the system through RTMP.
    - The system processes the RTMP stream.

2. **HLS Creation and Delivery**:
    - The RTMP stream is processed by FFmpeg to generate HLS segments and playlists.
    - The output is stored and served to viewers.

3. **Multi-Resolution Handling**:
    - Streams are processed based on predefined resolutions and bitrates (e.g., 720p, 480p).
    - Supports adaptive streaming to cater to different devices and internet speeds.

4. **Frontend Control Panel**:
    - Users and admins can monitor stream statistics, control live streams, and configure options.

5. **Database and Caching**:
    - **MySQL**: Stores persistent data such as user info, stream details, and logs.
    - **Redis**: Provides fast, temporary storage for frequently accessed settings (like system configurations).

---

## Prerequisites

To run this project, ensure you have:

- **Docker**: Installed on your system (version 20+).
- **Golang**: Installed for development-level tweaks if needed.
- **MySQL**: Used as the primary database.
- **Redis**: For caching.
- **FFmpeg**: Installed for handling video encoding tasks.

---

## Project Setup

### 1. Clone the Repository
```bash
git clone https://github.com/your-repo-name/live-streaming-system.git
cd live-streaming-system
```

### 2. Configure Environment Variables

Create a `.env` file in the project root and define the following values:
```env
GIN_MODE="release"

S3_API_KEY=
S3_BUCKET_NAME=
S3_DOMAIN=
S3_REGION=
S3_SECRET=

DB_DSN=
RABBITMQ_DSN=

REDIS_DSN=

GIN_PORT=

GRPC_PORT=

GRPC_AUTH_ADDRESS=
GRPC_USER_ADDRESS=
GRPC_COMMUNICATION_ADDRESS=
GRPC_ANALYTIC_ADDRESS=
GRPC_RTMP_ADDRESS=
GRPC_HLSMUX_ADDRESS=
GRPC_VIDEO_ADDRESS=

JWT_SECRET=
```

### 3. Build and Run the Project

#### Using Docker:
1. Build the project:
   ```bash
   docker-compose build
   ```

2. Start the containers:
   ```bash
   docker-compose up
   ```

#### Without Docker:
1. Start Redis and MySQL services on your system.
2. Run the backend server:
   ```bash
   go run main.go
   ```
3. Serve the frontend app via your preferred JavaScript hosting solution.

---

## Usage

### Streaming Workflow:

1. **Stream Input**:
    - Use a tool or service like **OBS Studio** to push a live stream to the server via RTMP:
      ```bash
      rtmp://<server_ip>:1935/stream/<stream_key>
      ```

2. **Stream Output**:
    - The HLS output will be available at:
      ```
      http://<server_ip>:8080/static/<stream_id>/master.m3u8
      ```

3. Multi-resolution playback can be added using a video player such as **Video.js** that supports HLS.

---


#
## Contributing

We welcome contributions! Please fork the repository and submit a pull request with your changes.

---

## License

This project is licensed under the **MIT License**. Feel free to use and modify it.

---