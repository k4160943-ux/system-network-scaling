# Configuration

This directory contains configuration files for the System Network Scaling Lab.

## Files

- **default.yaml** - Default configuration with all available options documented
- **.env** - Environment-specific overrides (not checked in, copy from .env.example)
- **profiles/** - Environment-specific configurations (development, staging, production)

## Loading Configuration

Configuration is loaded in the following priority order (highest to lowest):

1. **Environment Variables** - Override all other sources
2. **.env File** - Local environment overrides
3. **default.yaml** - Application defaults
4. **Built-in Defaults** - Hardcoded in code

## Configuration Structure

```yaml
application:      # General app settings
ingress:          # External data ingestion
normalization:    # Data transformation
transport:        # Message transport
consumers:        # Data consumption
observability:    # Metrics, logging, tracing
performance:      # Resource tuning
security:         # Security policies
scaling:          # Scaling strategy
cloud:            # Cloud integration
```

## Environment Variables

Create a `.env` file from `.env.example`:

```bash
cp .env.example .env
# Edit .env with your environment-specific values
```

Environment variables override YAML configuration:
- Use `_` to denote nested paths: `INGRESS_PORT=9000`
- Use nested objects: `SECURITY_TLS_ENABLED=true`

## Local Development

For local development, use:

```bash
# Load default.yaml with environment overrides
export ENV=development
export LOG_LEVEL=DEBUG
python -m system_network_scaling.ingress
```

## Docker

When running in Docker, mount configuration:

```bash
docker-compose -f docker-compose.yml up
```

The container will load:
1. `config/default.yaml` as defaults
2. Environment variables from `docker-compose.yml`
3. Any mounted `.env` file

## Production

For production deployments:

1. Create `config/production.yaml` with production settings
2. Use secrets management (not checked-in .env files)
3. Set environment variables via your orchestration platform (K8s, ECS, etc.)
