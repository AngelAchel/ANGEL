# ANGEL API Documentation

## Authentication

All API requests require JWT authentication.

```bash
curl -H "Authorization: Bearer <JWT_TOKEN>" https://<domain>:8443/api/v1/...
```

## Endpoints

### Teamserver API (Port 8443)

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | /api/v1/status | Server status |
| GET | /api/v1/agents | List active agents |
| POST | /api/v1/agents | Register agent |
| DELETE | /api/v1/agents/:id | Remove agent |
| GET | /api/v1/commands | List commands |
| POST | /api/v1/commands | Execute command |
| GET | /api/v1/reports | Get reports |

### Console API (Port 3000)

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | /api/v1/dashboard | Dashboard data |
| GET | /api/v1/layers | Layer status |
| GET | /api/v1/config | Configuration |
| POST | /api/v1/config | Update configuration |

### Rules Engine (Port 9444)

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | /api/v1/rules | List rules |
| POST | /api/v1/rules | Add rule |
| DELETE | /api/v1/rules/:id | Remove rule |
| PUT | /api/v1/rules/:id | Update rule |

## Response Format

```json
{
  "status": "success",
  "data": {},
  "error": null
}
```

## Error Codes

| Code | Meaning |
|------|---------|
| 200 | Success |
| 400 | Bad Request |
| 401 | Unauthorized |
| 403 | Forbidden |
| 404 | Not Found |
| 500 | Internal Server Error |

## Rate Limiting

- 100 requests/minute per IP
- 1000 requests/hour per IP