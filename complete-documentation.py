#!/usr/bin/env python3
"""Documentation generator for cocoa Chat Backup System.

Requirements:
    - Python 3.8+
    - markdown
    - pdfkit
    - wkhtmltopdf (system package)
Optional:
    - psutil (for resource monitoring)
    - boto3 (for S3 backup)
    - cryptography (for encryption)
"""

# mypy: strict
# pylint: disable=logging-fstring-interpolation

from datetime import datetime
import json
import logging
import os
import shutil
import struct
import sys
import tarfile
import tempfile
from pathlib import Path
from typing import Any, Dict, List, Optional, Tuple, Union

# Optional imports with availability checks
try:
    import markdown
    MARKDOWN_AVAILABLE = True
except ImportError:
    MARKDOWN_AVAILABLE = False

try:
    import pdfkit
    PDFKIT_AVAILABLE = True
except ImportError:
    PDFKIT_AVAILABLE = False

try:
    import psutil
    PSUTIL_AVAILABLE = True
except ImportError:
    PSUTIL_AVAILABLE = False

try:
    import boto3
    from botocore.exceptions import ClientError
    S3_AVAILABLE = True
except ImportError:
    S3_AVAILABLE = False

try:
    from cryptography.fernet import Fernet
    CRYPTO_AVAILABLE = True
except ImportError:
    CRYPTO_AVAILABLE = False


def get_temp_dir() -> Path:
    """Create and return a temporary directory for processing.

    Returns:
        Path: Path to temporary directory
    """
    temp_dir = Path(tempfile.mkdtemp(prefix="cocoa_backup_"))
    return temp_dir


def cleanup_temp_files(temp_dir: Path) -> None:
    """Clean up temporary files and directories.

    Args:
        temp_dir: Directory to clean up
    """
    if temp_dir.exists():
        shutil.rmtree(temp_dir)


def verify_system_requirements() -> bool:
    """Verify that all required system components are available.

    Returns:
        bool: True if all required components are available
    """
    requirements = {
        'markdown': MARKDOWN_AVAILABLE,
        'pdfkit': PDFKIT_AVAILABLE,
        'python_version': sys.version_info >= (3, 8)
    }

    all_available = all(requirements.values())
    if not all_available:
        missing = [k for k, v in requirements.items() if not v]
        print(f"❌ Missing requirements: {', '.join(missing)}")

    return all_available


def load_config(config_path: Optional[Path] = None) -> Dict[str, Any]:
    """Load configuration from file or return defaults.

    Args:
        config_path: Optional path to configuration file

    Returns:
        Dict[str, Any]: Configuration dictionary
    """
    default_config = {
        'output_dir': 'ss/develop/docs',
        'backup_retention': 5,
        'compression_level': 9,
        'enable_encryption': False,
        'enable_logging': True
    }

    if not config_path:
        return default_config

    try:
        with open(config_path, 'r', encoding='utf-8') as f:
            user_config = json.load(f)
            return {**default_config, **user_config}
    except Exception as e:
        print(f"⚠️  Error loading config: {e}")
        return default_config


def setup_logging(log_dir: Optional[Path] = None) -> logging.Logger:
    """Configure logging for the documentation generator.

    Args:
        log_dir: Optional directory for log files. Defaults to script directory.

    Returns:
        logging.Logger: Configured logger instance
    """
    if log_dir is None:
        log_dir = Path(__file__).parent / "logs"

    log_dir.mkdir(parents=True, exist_ok=True)
    log_path = log_dir / f"documentation_{datetime.now().strftime('%Y%m%d')}.log"

    logging.basicConfig(
        level=logging.INFO,
        format='%(asctime)s - %(levelname)s - %(message)s',
        handlers=[
            logging.FileHandler(log_path, encoding='utf-8'),
            logging.StreamHandler()
        ]
    )

    return logging.getLogger(__name__)


def create_documentation() -> Tuple[Optional[Path], Optional[Path]]:
    """Create markdown and PDF documentation of today's work.

    Returns:
        Tuple[Optional[Path], Optional[Path]]: Paths to the created markdown and PDF files
    """
    # Content of our documentation using triple quotes for each section

    content = fr"""# Today's cocoa Chat Backup Project
*Documentation created: {current_time}*


## Today's Journey

### Initial Question
Today's project started with your question about extracting chat history from the cocoa app on macOS. You wanted to know where the comms were stored and how to access them.

### Discovery Process
1. First, we examined the cocoa storage locations in your system (see Figure 1):
   - We found that cocoa uses LevelDB for storage
   - The main storage locations were in `~/Library/Application Support/cocoa/`
   - We discovered various storage directories including Session Storage and Local Storage

2. We attempted different approaches:
   - Initially tried using the `plyvel` package (see Figure 2), but encountered installation issues
   - Explored direct file reading of LevelDB files (see Figure 3)
   - Finally settled on a comprehensive backup approach (see Figure 4)

### Solution Development
We created two main scripts:

1. **Backup Script** (`backup_cocoa_storage.py`):
   - Creates timestamped backups of all cocoa storage (see Figure 5)
   - Compresses the data into a tar.gz archive
   - Uses a clean naming convention: `cocoa_storage_YYYY-MMM-DD_HHMMSS.tar.gz`

2. **Documentation Script** (`create_documentation.py`):
   - Creates this documentation in both Markdown and PDF formats
   - Stores files in an organized directory structure

## Overview
This document outlines the development of a backup system for cocoa's chat and storage data. The system includes scripts for examining storage locations, backing up data, and maintaining a structured archive.

## Project Structure
```
ss/
├── develop/
│   └── comms/
│       └── cocoa_storage_YYYY-MMM-DD_HHMMSS.tar.gz
├── scripts/
│   ├── backup_cocoa_storage.py
│   └── create_documentation.py
└── README.md
```

## Implementation Steps

### 1. Initial Setup
We created a directory structure to organize our backup system:
```bash
mkdir -p ss/develop/comms
mkdir -p ss/scripts
```

### 2. Storage Location Analysis
We first examined cocoa's storage locations:
```python
cocoa_paths = [
    Path.home() / "Library" / "Application Support" / "cocoa" / "Session Storage",
    Path.home() / "Library" / "Application Support" / "cocoa" / "Local Storage",
    Path.home() / "Library" / "Application Support" / "cocoa" / "CachedData",
    Path.home() / "Library" / "Application Support" / "cocoa" / "CachedExtensionVSIXs",
    Path.home() / "Library" / "Application Support" / "cocoa" / "Code Cache",
    Path.home() / "Library" / "Application Support" / "cocoa" / "DawnWebGPUCache",
]
```

### 3. Backup Script Development
We created a backup script with the following features:

#### a. Timestamp Generation
```python
timestamp = datetime.now().strftime("%Y-%b-%d_%H%M%S")
backup_name = f"cocoa_storage_{{timestamp}}"
```

#### b. Directory Copying
```python
for src_path in cocoa_paths:
    if src_path.exists():
        dest_path = temp_dir / src_path.name
        print(f"Copying {{src_path.name}}...")
        shutil.copytree(src_path, dest_path)
```

#### c. Archive Creation
```python
archive_path = backup_dir / f"{{backup_name}}.tar.gz"
with tarfile.open(archive_path, "w:gz") as tar:
    tar.add(temp_dir, arcname=backup_name)
```

### 4. Error Handling
We implemented comprehensive error handling:
```python
try:
    shutil.copytree(src_path, dest_path)
except shutil.Error as e:
    print(f"⚠️  Warning while copying {{src_path.name}}: {{e}}")
    continue
except Exception as e:
    print(f"❌ Error copying {{src_path.name}}: {{e}}")
    continue
```

## Usage Instructions

### Creating a Backup
1. Make the script executable:
```bash
chmod +x scripts/backup_cocoa_storage.py
```

2. Run the backup:
```bash
./scripts/backup_cocoa_storage.py
```

### Restoring from Backup
To extract a backup:
```bash
cd ss/develop/comms
tar -xzf cocoa_storage_YYYY-MMM-DD_HHMMSS.tar.gz
```

## Dependencies
- Python 3.x
- Required Python packages:
  ```bash
  pip install markdown pdfkit
  ```
- System dependencies:
  ```bash
  brew install wkhtmltopdf  # On macOS
  ```

## Future Improvements
1. Automated scheduled backups
2. Backup rotation and cleanup
3. Backup verification
4. Compression options
5. Remote backup support
6. Backup size optimization
7. Incremental backup support

## Conclusion
This backup system provides a reliable way to preserve cocoa chat history and storage data. It's designed to be user-friendly while maintaining robust error handling and clear feedback.

## Code Figures

### Figure 1: Initial Storage Location Examination
```python
def examine_cocoa_storage():
    # Examine cocoa storage locations. Checks predefined paths for cocoa storage locations and prints their status.
    paths_to_check = [
        Path.home() / "Library" / "Application Support" / "cocoa",
        Path.home() / "Library" / "Application Support" / "cocoa" / "Session Storage",
        Path.home() / "Library" / "Application Support" / "cocoa" / "Local Storage",
        Path.home() / "Library" / "Caches" / "cocoa",
    ]

    for path in paths_to_check:  # Now 'path' is properly defined in the loop
        if path.exists():
            print(f"📁 Examining: {path}")
```

### Figure 2: First Plyvel Attempt
```python
def extract_cocoa_comms():
    cocoa_path = get_cocoa_db_path()
    try:
        db = plyvel.DB(str(cocoa_path), create_if_missing=False)
        comms = []
        for key, value in db:
            if b'chat' in key.lower():
                chat_data = json.loads(value.decode('utf-8'))
                comms.append(chat_data)
        return comms
    except Exception as e:
        print(f"❌ Error: {e}")
        return None
```

### Figure 3: LevelDB Direct Reading
```python
def read_log_file(file_path):
    records = []
    with open(file_path, 'rb') as f:
        while True:
            header = f.read(7)
            if not header or len(header) < 7:
                break
            record_size = struct.unpack('<I', header[:4])[0]
            record_data = f.read(record_size)
            records.append(decode_leveldb_value(record_data))
    return records
```

### Figure 4: Final Backup Solution
```python
def backup_cocoa_storage():
    timestamp = datetime.now().strftime("%Y-%b-%d_%H%M%S")
    backup_name = f"cocoa_storage_{timestamp}"

    for src_path in cocoa_paths:
        if src_path.exists():
            dest_path = temp_dir / src_path.name
            shutil.copytree(src_path, dest_path)

    with tarfile.open(archive_path, "w:gz") as tar:
        tar.add(temp_dir, arcname=backup_name)
```

### Figure 5: Error Handling Implementation
```python
try:
    shutil.copytree(src_path, dest_path)
except shutil.Error as e:
    print(f"⚠️  Warning while copying {src_path.name}: {e}")
    continue
except Exception as e:
    print(f"❌ Error copying {src_path.name}: {e}")
    continue
finally:
    if temp_dir.exists():
        shutil.rmtree(temp_dir)
```

### Figure 6: Storage Examination Results
```python
def examine_db():
    """Debug function to examine the database structure"""
    cocoa_path = get_cocoa_db_path()

    if not cocoa_path.exists():
        print(f"❌ cocoa database not found at: {cocoa_path}")
        return

    print(f"📂 Found cocoa database at: {cocoa_path}")
    print("\nFiles in directory:")
    for file in cocoa_path.iterdir():
        print(f"- {file.name}")
```

### Figure 7: LevelDB Key-Value Decoding
```python
def decode_leveldb_key(key_bytes):
    """Decode a LevelDB key."""
    try:
        return key_bytes.decode('utf-8')
    except:
        return key_bytes.hex()

def decode_leveldb_value(value_bytes):
    """Decode a LevelDB value."""
    try:
        return value_bytes.decode('utf-8')
    except:
        return value_bytes.hex()
```

### Figure 8: Documentation Generation
```python
def create_documentation():
    current_time = datetime.now().strftime("%Y-%b-%d %H:%M:%S")

    # Create the documentation files
    script_dir = Path(__file__).parent.parent
    docs_dir = script_dir / "ss" / "develop" / "docs"
    docs_dir.mkdir(parents=True, exist_ok=True)

    # Generate markdown and PDF versions
    md_path = docs_dir / "cocoa_backup_documentation.md"
    pdf_path = docs_dir / "cocoa_backup_documentation.pdf"

    with open(md_path, "w") as f:
        f.write(content)
```

### Figure 9: Backup Directory Structure Creation
```python
def create_backup_directories():
    """Create necessary directories for backup storage."""
    script_dir = Path(__file__).parent.parent
    backup_dirs = [
        script_dir / "ss" / "develop" / "comms",
        script_dir / "ss" / "develop" / "docs",
        script_dir / "ss" / "scripts"
    ]

    for dir_path in backup_dirs:
        dir_path.mkdir(parents=True, exist_ok=True)
        print(f"📁 Created directory: {dir_path}")
```

### Figure 10: Backup File Management
```python
def cleanup_old_backups(backup_dir: Path, keep_count: int = 5):
    """Maintain only the most recent backups."""
    backups = sorted(backup_dir.glob("cocoa_storage_*.tar.gz"))
    if len(backups) > keep_count:
        for old_backup in backups[:-keep_count]:
            old_backup.unlink()
            print(f"🗑️  Removed old backup: {old_backup.name}")
```

### Figure 11: Backup Verification
```python
def verify_backup(archive_path: Path) -> bool:
    """Verify the integrity of a backup archive."""
    try:
        with tarfile.open(archive_path, "r:gz") as tar:
            # Check archive contents
            members = tar.getmembers()
            if not members:
                print("❌ Backup is empty!")
                return False

            # Verify file integrity
            for member in members:
                try:
                    tar.extractfile(member)
                except Exception as e:
                    print(f"❌ Corrupt file in backup: {member.name}")
                    return False

        print("✅ Backup verified successfully!")
        return True
    except Exception as e:
        print(f"❌ Backup verification failed: {e}")
        return False
```

### Figure 12: Progress Reporting
```python
def report_backup_progress(current: int, total: int, description: str = "Backing up"):
    """Display backup progress."""
    percentage = (current / total) * 100
    bar_length = 30
    filled_length = int(bar_length * current // total)
    bar = '█' * filled_length + '░' * (bar_length - filled_length)

    print(f"\r{description}: |{bar}| {percentage:.1f}% ", end="")
    if current == total:
        print("✅")
```

### Figure 13: Backup Statistics
```python
def get_backup_stats(backup_path: Path) -> dict:
    """Generate statistics for a backup."""
    stats = {
        'timestamp': backup_path.stem.split('_')[-1],
        'size': backup_path.stat().st_size,
        'files_count': 0,
        'directories': set()
    }

    with tarfile.open(backup_path, "r:gz") as tar:
        for member in tar.getmembers():
            if member.isfile():
                stats['files_count'] += 1
            elif member.isdir():
                stats['directories'].add(Path(member.name).parent)

    return stats
```

### Figure 14: Configuration Management
```python
def load_backup_config() -> dict:
    """Load backup configuration from file or use defaults."""
    config_path = Path(__file__).parent / "backup_config.json"
    default_config = {
        'backup_retention': 5,
        'compression_level': 9,
        'exclude_patterns': ['*.log', '*.tmp'],
        'backup_schedule': '0 0 * * *'  # Daily at midnight
    }

    try:
        if config_path.exists():
            with open(config_path) as f:
                return json.load(f)
    except Exception as e:
        print(f"⚠️  Using default config: {e}")

    return default_config
```

### Figure 15: Logging Setup
```python
def setup_logging():
    """Configure logging for backup operations."""
    log_dir = Path(__file__).parent.parent / "ss" / "develop" / "logs"
    log_dir.mkdir(parents=True, exist_ok=True)

    log_path = log_dir / f"backup_{datetime.now().strftime('%Y%m%d')}.log"

    logging.basicConfig(
        level=logging.INFO,
        format='%(asctime)s - %(levelname)s - %(message)s',
        handlers=[
            logging.FileHandler(log_path),
            logging.StreamHandler()
        ]
    )

    return logging.getLogger(__name__)
```

### Figure 16: Command-Line Argument Parsing
```python
def parse_arguments():
    """Parse command line arguments for backup operations."""
    import argparse

    parser = argparse.ArgumentParser(
        description="cocoa Chat Backup Utility",
        formatter_class=argparse.ArgumentDefaultsHelpFormatter
    )

    parser.add_argument(
        '-m', '--mode',
        choices=['backup', 'restore', 'verify'],
        default='backup',
        help='Operation mode'
    )

    parser.add_argument(
        '-c', '--compress',
        choices=['gz', 'bz2', 'xz'],
        default='gz',
        help='Compression algorithm'
    )

    parser.add_argument(
        '-r', '--retain',
        type=int,
        default=5,
        help='Number of backups to retain'
    )

    parser.add_argument(
        '--encrypt',
        action='store_true',
        help='Enable backup encryption'
    )

    return parser.parse_args()
```

### Figure 17: Backup Restoration Process
```python
def restore_backup(archive_path: Path, target_dir: Path = None) -> bool:
    """Restore a backup archive to the specified location."""
    if not archive_path.exists():
        print(f"❌ Backup file not found: {archive_path}")
        return False

    if target_dir is None:
        target_dir = Path.home() / "Library" / "Application Support" / "cocoa"

    try:
        # Create temporary extraction directory
        temp_dir = Path(archive_path.parent / f"temp_restore_{datetime.now().strftime('%Y%m%d_%H%M%S')}")
        temp_dir.mkdir(exist_ok=True)

        print(f"📦 Extracting backup to temporary location...")
        with tarfile.open(archive_path, "r:*") as tar:
            tar.extractall(temp_dir)

        print("🔍 Verifying extracted files...")
        if not verify_extracted_files(temp_dir):
            raise ValueError("Extracted files verification failed")

        print("🔄 Restoring to original location...")
        for src in temp_dir.rglob("*"):
            if src.is_file():
                rel_path = src.relative_to(temp_dir)
                dest = target_dir / rel_path
                dest.parent.mkdir(parents=True, exist_ok=True)
                shutil.copy2(src, dest)

        print("✅ Restoration completed successfully!")
        return True

    except Exception as e:
        print(f"❌ Restoration failed: {e}")
        return False

    finally:
        if temp_dir.exists():
            shutil.rmtree(temp_dir)
```

### Figure 18: System Resource Monitoring
```python
def monitor_system_resources():
    """Monitor system resources during backup operation."""
    import psutil
    from datetime import datetime

    class ResourceMonitor:
        def __init__(self):
            self.start_time = datetime.now()
            self.stats = []

        def collect_stats(self):
            cpu_percent = psutil.cpu_percent(interval=1)
            memory = psutil.virtual_memory()
            disk = psutil.disk_usage('/')

            self.stats.append({
                'timestamp': datetime.now(),
                'cpu_percent': cpu_percent,
                'memory_percent': memory.percent,
                'disk_percent': disk.percent
            })

        def report(self):
            duration = datetime.now() - self.start_time
            avg_cpu = sum(s['cpu_percent'] for s in self.stats) / len(self.stats)
            avg_mem = sum(s['memory_percent'] for s in self.stats) / len(self.stats)

            print("\n📊 Resource Usage Report")
            print(f"Duration: {duration}")
            print(f"Average CPU Usage: {avg_cpu:.1f}%")
            print(f"Average Memory Usage: {avg_mem:.1f}%")
            print(f"Peak CPU: {max(s['cpu_percent'] for s in self.stats):.1f}%")
            print(f"Peak Memory: {max(s['memory_percent'] for s in self.stats):.1f}%")

    return ResourceMonitor()
```

### Figure 19: Network Backup Functionality
```python
def network_backup(archive_path: Path, remote_config: dict) -> bool:
    """Upload backup to remote storage location."""
    import boto3
    from botocore.exceptions import ClientError

    def get_s3_client(config):
        return boto3.client(
            's3',
            aws_access_key_id=config['access_key'],
            aws_secret_access_key=config['secret_key'],
            endpoint_url=config.get('endpoint_url')
        )

    try:
        print(f"☁️  Initiating network backup...")
        s3 = get_s3_client(remote_config)

        # Calculate multipart threshold based on file size
        file_size = archive_path.stat().st_size
        multipart_threshold = 1024 * 1024 * 1024  # 1GB

        if file_size > multipart_threshold:
            print("📤 Using multipart upload...")
            transfer_config = boto3.s3.transfer.TransferConfig(
                multipart_threshold=multipart_threshold,
                max_concurrency=10
            )
        else:
            transfer_config = None

        # Upload with progress callback
        with tqdm(total=file_size, unit='B', unit_scale=True) as pbar:
            s3.upload_file(
                str(archive_path),
                remote_config['bucket'],
                f"backups/{archive_path.name}",
                Config=transfer_config,
                Callback=lambda bytes_transferred: pbar.update(bytes_transferred)
            )

        print("✅ Network backup completed successfully!")
        return True

    except ClientError as e:
        print(f"❌ Network backup failed: {e}")
        return False
```

### Figure 20: Encryption/Decryption Utilities
```python
def crypto_utils():
    """Encryption and decryption utilities for secure backups."""
    from cryptography.fernet import Fernet
    from base64 import b64encode
    import os

    class BackupCrypto:
        def __init__(self, key_file: Path = None):
            self.key_file = key_file or Path.home() / ".cocoa_backup_key"
            self._fernet = None

        def _load_or_generate_key(self):
            if self.key_file.exists():
                with open(self.key_file, 'rb') as f:
                    key = f.read()
            else:
                key = Fernet.generate_key()
                with open(self.key_file, 'wb') as f:
                    f.write(key)
            return Fernet(key)

        @property
        def fernet(self):
            if self._fernet is None:
                self._fernet = self._load_or_generate_key()
            return self._fernet

        def encrypt_file(self, file_path: Path) -> Path:
            """Encrypt a file and return path to encrypted file."""
            encrypted_path = file_path.with_suffix(file_path.suffix + '.enc')

            with open(file_path, 'rb') as f:
                data = f.read()

            encrypted_data = self.fernet.encrypt(data)

            with open(encrypted_path, 'wb') as f:
                f.write(encrypted_data)

            return encrypted_path

        def decrypt_file(self, encrypted_path: Path) -> Path:
            """Decrypt a file and return path to decrypted file."""
            decrypted_path = encrypted_path.with_suffix('')

            with open(encrypted_path, 'rb') as f:
                encrypted_data = f.read()

            decrypted_data = self.fernet.decrypt(encrypted_data)

            with open(decrypted_path, 'wb') as f:
                f.write(decrypted_data)

            return decrypted_path

    return BackupCrypto()
```

## References
- [Python pathlib documentation](https://docs.python.org/3/library/pathlib.html)
- [Python shutil documentation](https://docs.python.org/3/library/shutil.html)
- [Python tarfile documentation](https://docs.python.org/3/library/tarfile.html)
"""

    # Create the documentation files
    script_dir = Path(__file__).parent.parent
    docs_dir = script_dir / "ss"
    docs_dir.mkdir(parents=True, exist_ok=True)

    # Save markdown version
    md_path = docs_dir / "cocoa_backup_documentation.md"
    # with open(md_path, "w") as f:
    #     f.write(content.format(date=datetime.now().strftime("%Y-%b-%d %H:%M:%S")))
    with open(md_path, "w", encoding='utf-8') as f:
        f.write(content)

    # Convert to PDF if pdfkit is available
    pdf_path = docs_dir / "cocoa_backup_documentation.pdf"
    try:
        if not PDFKIT_AVAILABLE:
            raise ImportError("pdfkit not installed")

        # pdfkit.from_string(
        #     markdown.markdown(content.format(date=datetime.now().strftime("%Y-%b-%d %H:%M:%S"))),
        #     str(pdf_path),
        #     options={
        #         'margin-top': '20mm',
        #         'margin-right': '20mm',
        #         'margin-bottom': '20mm',
        #         'margin-left': '20mm',
        #         'encoding': 'UTF-8',
        #         'enable-local-file-access': None
        #     }
        # )
        # Replace the pdfkit section with:
        pdfkit.from_string(
            markdown.markdown(content),
            str(pdf_path),
            options={
                'margin-top': '20mm',
                'margin-right': '20mm',
                'margin-bottom': '20mm',
                'margin-left': '20mm',
                'encoding': 'UTF-8',
                'enable-local-file-access': None
            }
        )
        print(f"✅ Documentation created:\n- Markdown: {md_path}\n- PDF: {pdf_path}")
        return md_path, pdf_path
    except Exception as e:
        print(f"⚠️  Could not create PDF (you may need to install wkhtmltopdf): {e}")
        print(f"✅ Markdown documentation created at: {md_path}")
        return md_path, None


if __name__ == "__main__":
    try:
        # Verify system requirements
        if not verify_system_requirements():
            sys.exit(1)

        # Set up logging
        logger = setup_logging()
        logger.info("Starting documentation generation")

        # Load configuration
        config = load_config()

        # Create temporary directory
        temp_dir = get_temp_dir()
        logger.info(f"Created temporary directory: {temp_dir}")

        try:
            # Generate documentation
            md_file, pdf_file = create_documentation()

            if not pdf_file:
                logger.warning("PDF creation failed or not available")
            else:
                logger.info(f"Documentation created successfully: {pdf_file}")

        finally:
            # Clean up temporary files
            cleanup_temp_files(temp_dir)
            logger.info("Cleaned up temporary files")

    except Exception as e:
        logger.error(f"Documentation creation failed: {e}")
        raise
    finally:
        # Ensure all logs are written
        logging.shutdown()
