// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new DB instance. The new DB instance can be an RDS DB instance, or it
// can be a DB instance in an Aurora DB cluster. For an Aurora DB cluster, you can
// call this operation multiple times to add more than one DB instance to the
// cluster. For more information about creating an RDS DB instance, see Creating
// an Amazon RDS DB instance (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_CreateDBInstance.html)
// in the Amazon RDS User Guide. For more information about creating a DB instance
// in an Aurora DB cluster, see Creating an Amazon Aurora DB cluster (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Aurora.CreateInstance.html)
// in the Amazon Aurora User Guide.
func (c *Client) CreateDBInstance(ctx context.Context, params *CreateDBInstanceInput, optFns ...func(*Options)) (*CreateDBInstanceOutput, error) {
	if params == nil {
		params = &CreateDBInstanceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDBInstance", params, optFns, c.addOperationCreateDBInstanceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDBInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDBInstanceInput struct {

	// The compute and memory capacity of the DB instance, for example db.m5.large .
	// Not all DB instance classes are available in all Amazon Web Services Regions, or
	// for all database engines. For the full list of DB instance classes, and
	// availability for your engine, see DB instance classes (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.DBInstanceClass.html)
	// in the Amazon RDS User Guide or Aurora DB instance classes (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Concepts.DBInstanceClass.html)
	// in the Amazon Aurora User Guide.
	//
	// This member is required.
	DBInstanceClass *string

	// The identifier for this DB instance. This parameter is stored as a lowercase
	// string. Constraints:
	//   - Must contain from 1 to 63 letters, numbers, or hyphens.
	//   - First character must be a letter.
	//   - Can't end with a hyphen or contain two consecutive hyphens.
	// Example: mydbinstance
	//
	// This member is required.
	DBInstanceIdentifier *string

	// The database engine to use for this DB instance. Not every database engine is
	// available in every Amazon Web Services Region. Valid Values:
	//   - aurora-mysql (for Aurora MySQL DB instances)
	//   - aurora-postgresql (for Aurora PostgreSQL DB instances)
	//   - custom-oracle-ee (for RDS Custom for Oracle DB instances)
	//   - custom-oracle-ee-cdb (for RDS Custom for Oracle DB instances)
	//   - custom-sqlserver-ee (for RDS Custom for SQL Server DB instances)
	//   - custom-sqlserver-se (for RDS Custom for SQL Server DB instances)
	//   - custom-sqlserver-web (for RDS Custom for SQL Server DB instances)
	//   - mariadb
	//   - mysql
	//   - oracle-ee
	//   - oracle-ee-cdb
	//   - oracle-se2
	//   - oracle-se2-cdb
	//   - postgres
	//   - sqlserver-ee
	//   - sqlserver-se
	//   - sqlserver-ex
	//   - sqlserver-web
	//
	// This member is required.
	Engine *string

	// The amount of storage in gibibytes (GiB) to allocate for the DB instance. This
	// setting doesn't apply to Amazon Aurora DB instances. Aurora cluster volumes
	// automatically grow as the amount of data in your database increases, though you
	// are only charged for the space that you use in an Aurora cluster volume. Amazon
	// RDS Custom RDS for MariaDB RDS for MySQL RDS for Oracle RDS for PostgreSQL RDS
	// for SQL Server Constraints to the amount of storage for each storage type are
	// the following:
	//   - General Purpose (SSD) storage (gp2, gp3): Must be an integer from 40 to
	//   65536 for RDS Custom for Oracle, 16384 for RDS Custom for SQL Server.
	//   - Provisioned IOPS storage (io1): Must be an integer from 40 to 65536 for RDS
	//   Custom for Oracle, 16384 for RDS Custom for SQL Server.
	// Constraints to the amount of storage for each storage type are the following:
	//   - General Purpose (SSD) storage (gp2, gp3): Must be an integer from 20 to
	//   65536.
	//   - Provisioned IOPS storage (io1): Must be an integer from 100 to 65536.
	//   - Magnetic storage (standard): Must be an integer from 5 to 3072.
	// Constraints to the amount of storage for each storage type are the following:
	//   - General Purpose (SSD) storage (gp2, gp3): Must be an integer from 20 to
	//   65536.
	//   - Provisioned IOPS storage (io1): Must be an integer from 100 to 65536.
	//   - Magnetic storage (standard): Must be an integer from 5 to 3072.
	// Constraints to the amount of storage for each storage type are the following:
	//   - General Purpose (SSD) storage (gp2, gp3): Must be an integer from 20 to
	//   65536.
	//   - Provisioned IOPS storage (io1): Must be an integer from 100 to 65536.
	//   - Magnetic storage (standard): Must be an integer from 10 to 3072.
	// Constraints to the amount of storage for each storage type are the following:
	//   - General Purpose (SSD) storage (gp2, gp3): Must be an integer from 20 to
	//   65536.
	//   - Provisioned IOPS storage (io1): Must be an integer from 100 to 65536.
	//   - Magnetic storage (standard): Must be an integer from 5 to 3072.
	// Constraints to the amount of storage for each storage type are the following:
	//   - General Purpose (SSD) storage (gp2, gp3):
	//   - Enterprise and Standard editions: Must be an integer from 20 to 16384.
	//   - Web and Express editions: Must be an integer from 20 to 16384.
	//   - Provisioned IOPS storage (io1):
	//   - Enterprise and Standard editions: Must be an integer from 100 to 16384.
	//   - Web and Express editions: Must be an integer from 100 to 16384.
	//   - Magnetic storage (standard):
	//   - Enterprise and Standard editions: Must be an integer from 20 to 1024.
	//   - Web and Express editions: Must be an integer from 20 to 1024.
	AllocatedStorage *int32

	// Specifies whether minor engine upgrades are applied automatically to the DB
	// instance during the maintenance window. By default, minor engine upgrades are
	// applied automatically. If you create an RDS Custom DB instance, you must set
	// AutoMinorVersionUpgrade to false .
	AutoMinorVersionUpgrade *bool

	// The Availability Zone (AZ) where the database will be created. For information
	// on Amazon Web Services Regions and Availability Zones, see Regions and
	// Availability Zones (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.RegionsAndAvailabilityZones.html)
	// . For Amazon Aurora, each Aurora DB cluster hosts copies of its storage in three
	// separate Availability Zones. Specify one of these Availability Zones. Aurora
	// automatically chooses an appropriate Availability Zone if you don't specify one.
	// Default: A random, system-chosen Availability Zone in the endpoint's Amazon Web
	// Services Region. Constraints:
	//   - The AvailabilityZone parameter can't be specified if the DB instance is a
	//   Multi-AZ deployment.
	//   - The specified Availability Zone must be in the same Amazon Web Services
	//   Region as the current endpoint.
	// Example: us-east-1d
	AvailabilityZone *string

	// The number of days for which automated backups are retained. Setting this
	// parameter to a positive number enables backups. Setting this parameter to 0
	// disables automated backups. This setting doesn't apply to Amazon Aurora DB
	// instances. The retention period for automated backups is managed by the DB
	// cluster. Default: 1 Constraints:
	//   - Must be a value from 0 to 35.
	//   - Can't be set to 0 if the DB instance is a source to read replicas.
	//   - Can't be set to 0 for an RDS Custom for Oracle DB instance.
	BackupRetentionPeriod *int32

	// The location for storing automated backups and manual snapshots. Valie Values:
	//   - outposts (Amazon Web Services Outposts)
	//   - region (Amazon Web Services Region)
	// Default: region For more information, see Working with Amazon RDS on Amazon Web
	// Services Outposts (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-on-outposts.html)
	// in the Amazon RDS User Guide.
	BackupTarget *string

	// The CA certificate identifier to use for the DB instance's server certificate.
	// This setting doesn't apply to RDS Custom DB instances. For more information, see
	// Using SSL/TLS to encrypt a connection to a DB instance (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.SSL.html)
	// in the Amazon RDS User Guide and Using SSL/TLS to encrypt a connection to a DB
	// cluster (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/UsingWithRDS.SSL.html)
	// in the Amazon Aurora User Guide.
	CACertificateIdentifier *string

	// For supported engines, the character set ( CharacterSet ) to associate the DB
	// instance with. This setting doesn't apply to the following DB instances:
	//   - Amazon Aurora - The character set is managed by the DB cluster. For more
	//   information, see CreateDBCluster .
	//   - RDS Custom - However, if you need to change the character set, you can
	//   change it on the database itself.
	CharacterSetName *string

	// Spcifies whether to copy tags from the DB instance to snapshots of the DB
	// instance. By default, tags are not copied. This setting doesn't apply to Amazon
	// Aurora DB instances. Copying tags to snapshots is managed by the DB cluster.
	// Setting this value for an Aurora DB instance has no effect on the DB cluster
	// setting.
	CopyTagsToSnapshot *bool

	// The instance profile associated with the underlying Amazon EC2 instance of an
	// RDS Custom DB instance. This setting is required for RDS Custom. Constraints:
	//   - The profile must exist in your account.
	//   - The profile must have an IAM role that Amazon EC2 has permissions to
	//   assume.
	//   - The instance profile name and the associated IAM role name must start with
	//   the prefix AWSRDSCustom .
	// For the list of permissions required for the IAM role, see  Configure IAM and
	// your VPC (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-setup-orcl.html#custom-setup-orcl.iam-vpc)
	// in the Amazon RDS User Guide.
	CustomIamInstanceProfile *string

	// The identifier of the DB cluster that this DB instance will belong to. This
	// setting doesn't apply to RDS Custom DB instances.
	DBClusterIdentifier *string

	// The meaning of this parameter differs depending on the database engine. Amazon
	// Aurora MySQL Amazon Aurora PostgreSQL Amazon RDS Custom for Oracle Amazon RDS
	// Custom for SQL Server RDS for MariaDB RDS for MySQL RDS for Oracle RDS for
	// PostgreSQL RDS for SQL Server The name of the database to create when the
	// primary DB instance of the Aurora MySQL DB cluster is created. If you don't
	// specify a value, Amazon RDS doesn't create a database in the DB cluster.
	// Constraints:
	//   - Must contain 1 to 64 alphanumeric characters.
	//   - Can't be a word reserved by the database engine.
	// The name of the database to create when the primary DB instance of the Aurora
	// PostgreSQL DB cluster is created. Default: postgres Constraints:
	//   - Must contain 1 to 63 alphanumeric characters.
	//   - Must begin with a letter. Subsequent characters can be letters,
	//   underscores, or digits (0 to 9).
	//   - Can't be a word reserved by the database engine.
	// The Oracle System ID (SID) of the created RDS Custom DB instance. Default: ORCL
	// Constraints:
	//   - Must contain 1 to 8 alphanumeric characters.
	//   - Must contain a letter.
	//   - Can't be a word reserved by the database engine.
	// Not applicable. Must be null. The name of the database to create when the DB
	// instance is created. If you don't specify a value, Amazon RDS doesn't create a
	// database in the DB instance. Constraints:
	//   - Must contain 1 to 64 letters or numbers.
	//   - Must begin with a letter. Subsequent characters can be letters,
	//   underscores, or digits (0-9).
	//   - Can't be a word reserved by the database engine.
	// The name of the database to create when the DB instance is created. If you
	// don't specify a value, Amazon RDS doesn't create a database in the DB instance.
	// Constraints:
	//   - Must contain 1 to 64 letters or numbers.
	//   - Must begin with a letter. Subsequent characters can be letters,
	//   underscores, or digits (0-9).
	//   - Can't be a word reserved by the database engine.
	// The Oracle System ID (SID) of the created DB instance. Default: ORCL
	// Constraints:
	//   - Can't be longer than 8 characters.
	//   - Can't be a word reserved by the database engine, such as the string NULL .
	// The name of the database to create when the DB instance is created. Default:
	// postgres Constraints:
	//   - Must contain 1 to 63 letters, numbers, or underscores.
	//   - Must begin with a letter. Subsequent characters can be letters,
	//   underscores, or digits (0-9).
	//   - Can't be a word reserved by the database engine.
	// Not applicable. Must be null.
	DBName *string

	// The name of the DB parameter group to associate with this DB instance. If you
	// don't specify a value, then Amazon RDS uses the default DB parameter group for
	// the specified DB engine and version. This setting doesn't apply to RDS Custom DB
	// instances. Constraints:
	//   - Must be 1 to 255 letters, numbers, or hyphens.
	//   - The first character must be a letter.
	//   - Can't end with a hyphen or contain two consecutive hyphens.
	DBParameterGroupName *string

	// A list of DB security groups to associate with this DB instance. This setting
	// applies to the legacy EC2-Classic platform, which is no longer used to create
	// new DB instances. Use the VpcSecurityGroupIds setting instead.
	DBSecurityGroups []string

	// A DB subnet group to associate with this DB instance. Constraints:
	//   - Must match the name of an existing DB subnet group.
	//   - Must not be default .
	// Example: mydbsubnetgroup
	DBSubnetGroupName *string

	// Specifies whether the DB instance has deletion protection enabled. The database
	// can't be deleted when deletion protection is enabled. By default, deletion
	// protection isn't enabled. For more information, see Deleting a DB Instance (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_DeleteInstance.html)
	// . This setting doesn't apply to Amazon Aurora DB instances. You can enable or
	// disable deletion protection for the DB cluster. For more information, see
	// CreateDBCluster . DB instances in a DB cluster can be deleted even when deletion
	// protection is enabled for the DB cluster.
	DeletionProtection *bool

	// The Active Directory directory ID to create the DB instance in. Currently, only
	// Microsoft SQL Server, MySQL, Oracle, and PostgreSQL DB instances can be created
	// in an Active Directory Domain. For more information, see Kerberos Authentication (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/kerberos-authentication.html)
	// in the Amazon RDS User Guide. This setting doesn't apply to the following DB
	// instances:
	//   - Amazon Aurora (The domain is managed by the DB cluster.)
	//   - RDS Custom
	Domain *string

	// The name of the IAM role to use when making API calls to the Directory Service.
	// This setting doesn't apply to the following DB instances:
	//   - Amazon Aurora (The domain is managed by the DB cluster.)
	//   - RDS Custom
	DomainIAMRoleName *string

	// The list of log types that need to be enabled for exporting to CloudWatch Logs.
	// For more information, see Publishing Database Logs to Amazon CloudWatch Logs (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_LogAccess.html#USER_LogAccess.Procedural.UploadtoCloudWatch)
	// in the Amazon RDS User Guide. This setting doesn't apply to the following DB
	// instances:
	//   - Amazon Aurora (CloudWatch Logs exports are managed by the DB cluster.)
	//   - RDS Custom
	// The following values are valid for each DB engine:
	//   - RDS for MariaDB - audit | error | general | slowquery
	//   - RDS for Microsoft SQL Server - agent | error
	//   - RDS for MySQL - audit | error | general | slowquery
	//   - RDS for Oracle - alert | audit | listener | trace | oemagent
	//   - RDS for PostgreSQL - postgresql | upgrade
	EnableCloudwatchLogsExports []string

	// Specifies whether to enable a customer-owned IP address (CoIP) for an RDS on
	// Outposts DB instance. A CoIP provides local or external connectivity to
	// resources in your Outpost subnets through your on-premises network. For some use
	// cases, a CoIP can provide lower latency for connections to the DB instance from
	// outside of its virtual private cloud (VPC) on your local network. For more
	// information about RDS on Outposts, see Working with Amazon RDS on Amazon Web
	// Services Outposts (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-on-outposts.html)
	// in the Amazon RDS User Guide. For more information about CoIPs, see
	// Customer-owned IP addresses (https://docs.aws.amazon.com/outposts/latest/userguide/routing.html#ip-addressing)
	// in the Amazon Web Services Outposts User Guide.
	EnableCustomerOwnedIp *bool

	// Specifies whether to enable mapping of Amazon Web Services Identity and Access
	// Management (IAM) accounts to database accounts. By default, mapping isn't
	// enabled. For more information, see IAM Database Authentication for MySQL and
	// PostgreSQL (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.IAMDBAuth.html)
	// in the Amazon RDS User Guide. This setting doesn't apply to the following DB
	// instances:
	//   - Amazon Aurora (Mapping Amazon Web Services IAM accounts to database
	//   accounts is managed by the DB cluster.)
	//   - RDS Custom
	EnableIAMDatabaseAuthentication *bool

	// Specifies whether to enable Performance Insights for the DB instance. For more
	// information, see Using Amazon Performance Insights (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PerfInsights.html)
	// in the Amazon RDS User Guide. This setting doesn't apply to RDS Custom DB
	// instances.
	EnablePerformanceInsights *bool

	// The version number of the database engine to use. This setting doesn't apply to
	// Amazon Aurora DB instances. The version number of the database engine the DB
	// instance uses is managed by the DB cluster. For a list of valid engine versions,
	// use the DescribeDBEngineVersions operation. The following are the database
	// engines and links to information about the major and minor versions that are
	// available with Amazon RDS. Not every database engine is available for every
	// Amazon Web Services Region. Amazon RDS Custom for Oracle Amazon RDS Custom for
	// SQL Server RDS for MariaDB RDS for Microsoft SQL Server RDS for MySQL RDS for
	// Oracle RDS for PostgreSQL A custom engine version (CEV) that you have previously
	// created. This setting is required for RDS Custom for Oracle. The CEV name has
	// the following format: 19.customized_string. A valid CEV name is 19.my_cev1 . For
	// more information, see Creating an RDS Custom for Oracle DB instance (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-creating.html#custom-creating.create)
	// in the Amazon RDS User Guide. See RDS Custom for SQL Server general requirements (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-reqs-limits-MS.html)
	// in the Amazon RDS User Guide. For information, see MariaDB on Amazon RDS
	// versions (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_MariaDB.html#MariaDB.Concepts.VersionMgmt)
	// in the Amazon RDS User Guide. For information, see Microsoft SQL Server
	// versions on Amazon RDS (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_SQLServer.html#SQLServer.Concepts.General.VersionSupport)
	// in the Amazon RDS User Guide. For information, see MySQL on Amazon RDS versions (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_MySQL.html#MySQL.Concepts.VersionMgmt)
	// in the Amazon RDS User Guide. For information, see Oracle Database Engine
	// release notes (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.Oracle.PatchComposition.html)
	// in the Amazon RDS User Guide. For information, see Amazon RDS for PostgreSQL
	// versions and extensions (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_PostgreSQL.html#PostgreSQL.Concepts)
	// in the Amazon RDS User Guide.
	EngineVersion *string

	// The amount of Provisioned IOPS (input/output operations per second) to
	// initially allocate for the DB instance. For information about valid IOPS values,
	// see Amazon RDS DB instance storage (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Storage.html)
	// in the Amazon RDS User Guide. This setting doesn't apply to Amazon Aurora DB
	// instances. Storage is managed by the DB cluster. Constraints:
	//   - For RDS for MariaDB, MySQL, Oracle, and PostgreSQL - Must be a multiple
	//   between .5 and 50 of the storage amount for the DB instance.
	//   - For RDS for SQL Server - Must be a multiple between 1 and 50 of the storage
	//   amount for the DB instance.
	Iops *int32

	// The Amazon Web Services KMS key identifier for an encrypted DB instance. The
	// Amazon Web Services KMS key identifier is the key ARN, key ID, alias ARN, or
	// alias name for the KMS key. To use a KMS key in a different Amazon Web Services
	// account, specify the key ARN or alias ARN. This setting doesn't apply to Amazon
	// Aurora DB instances. The Amazon Web Services KMS key identifier is managed by
	// the DB cluster. For more information, see CreateDBCluster . If StorageEncrypted
	// is enabled, and you do not specify a value for the KmsKeyId parameter, then
	// Amazon RDS uses your default KMS key. There is a default KMS key for your Amazon
	// Web Services account. Your Amazon Web Services account has a different default
	// KMS key for each Amazon Web Services Region. For Amazon RDS Custom, a KMS key is
	// required for DB instances. For most RDS engines, if you leave this parameter
	// empty while enabling StorageEncrypted , the engine uses the default KMS key.
	// However, RDS Custom doesn't use the default key when this parameter is empty.
	// You must explicitly specify a key.
	KmsKeyId *string

	// The license model information for this DB instance. This setting doesn't apply
	// to Amazon Aurora or RDS Custom DB instances. Valid Values:
	//   - RDS for MariaDB - general-public-license
	//   - RDS for Microsoft SQL Server - license-included
	//   - RDS for MySQL - general-public-license
	//   - RDS for Oracle - bring-your-own-license | license-included
	//   - RDS for PostgreSQL - postgresql-license
	LicenseModel *string

	// Specifies whether to manage the master user password with Amazon Web Services
	// Secrets Manager. For more information, see Password management with Amazon Web
	// Services Secrets Manager (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-secrets-manager.html)
	// in the Amazon RDS User Guide. Constraints:
	//   - Can't manage the master user password with Amazon Web Services Secrets
	//   Manager if MasterUserPassword is specified.
	ManageMasterUserPassword *bool

	// The password for the master user. This setting doesn't apply to Amazon Aurora
	// DB instances. The password for the master user is managed by the DB cluster.
	// Constraints:
	//   - Can't be specified if ManageMasterUserPassword is turned on.
	//   - Can include any printable ASCII character except "/", """, or "@".
	// Length Constraints:
	//   - RDS for MariaDB - Must contain from 8 to 41 characters.
	//   - RDS for Microsoft SQL Server - Must contain from 8 to 128 characters.
	//   - RDS for MySQL - Must contain from 8 to 41 characters.
	//   - RDS for Oracle - Must contain from 8 to 30 characters.
	//   - RDS for PostgreSQL - Must contain from 8 to 128 characters.
	MasterUserPassword *string

	// The Amazon Web Services KMS key identifier to encrypt a secret that is
	// automatically generated and managed in Amazon Web Services Secrets Manager. This
	// setting is valid only if the master user password is managed by RDS in Amazon
	// Web Services Secrets Manager for the DB instance. The Amazon Web Services KMS
	// key identifier is the key ARN, key ID, alias ARN, or alias name for the KMS key.
	// To use a KMS key in a different Amazon Web Services account, specify the key ARN
	// or alias ARN. If you don't specify MasterUserSecretKmsKeyId , then the
	// aws/secretsmanager KMS key is used to encrypt the secret. If the secret is in a
	// different Amazon Web Services account, then you can't use the aws/secretsmanager
	// KMS key to encrypt the secret, and you must use a customer managed KMS key.
	// There is a default KMS key for your Amazon Web Services account. Your Amazon Web
	// Services account has a different default KMS key for each Amazon Web Services
	// Region.
	MasterUserSecretKmsKeyId *string

	// The name for the master user. This setting doesn't apply to Amazon Aurora DB
	// instances. The name for the master user is managed by the DB cluster. This
	// setting is required for RDS DB instances. Constraints:
	//   - Must be 1 to 16 letters, numbers, or underscores.
	//   - First character must be a letter.
	//   - Can't be a reserved word for the chosen database engine.
	MasterUsername *string

	// The upper limit in gibibytes (GiB) to which Amazon RDS can automatically scale
	// the storage of the DB instance. For more information about this setting,
	// including limitations that apply to it, see Managing capacity automatically
	// with Amazon RDS storage autoscaling (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PIOPS.StorageTypes.html#USER_PIOPS.Autoscaling)
	// in the Amazon RDS User Guide. This setting doesn't apply to the following DB
	// instances:
	//   - Amazon Aurora (Storage is managed by the DB cluster.)
	//   - RDS Custom
	MaxAllocatedStorage *int32

	// The interval, in seconds, between points when Enhanced Monitoring metrics are
	// collected for the DB instance. To disable collection of Enhanced Monitoring
	// metrics, specify 0 . If MonitoringRoleArn is specified, then you must set
	// MonitoringInterval to a value other than 0 . This setting doesn't apply to RDS
	// Custom DB instances. Valid Values: 0 | 1 | 5 | 10 | 15 | 30 | 60 Default: 0
	MonitoringInterval *int32

	// The ARN for the IAM role that permits RDS to send enhanced monitoring metrics
	// to Amazon CloudWatch Logs. For example, arn:aws:iam:123456789012:role/emaccess .
	// For information on creating a monitoring role, see Setting Up and Enabling
	// Enhanced Monitoring (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Monitoring.OS.html#USER_Monitoring.OS.Enabling)
	// in the Amazon RDS User Guide. If MonitoringInterval is set to a value other
	// than 0 , then you must supply a MonitoringRoleArn value. This setting doesn't
	// apply to RDS Custom DB instances.
	MonitoringRoleArn *string

	// Specifies whether the DB instance is a Multi-AZ deployment. You can't set the
	// AvailabilityZone parameter if the DB instance is a Multi-AZ deployment. This
	// setting doesn't apply to the following DB instances:
	//   - Amazon Aurora (DB instance Availability Zones (AZs) are managed by the DB
	//   cluster.)
	//   - RDS Custom
	MultiAZ *bool

	// The name of the NCHAR character set for the Oracle DB instance. This setting
	// doesn't apply to RDS Custom DB instances.
	NcharCharacterSetName *string

	// The network type of the DB instance. The network type is determined by the
	// DBSubnetGroup specified for the DB instance. A DBSubnetGroup can support only
	// the IPv4 protocol or the IPv4 and the IPv6 protocols ( DUAL ). For more
	// information, see Working with a DB instance in a VPC (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_VPC.WorkingWithRDSInstanceinaVPC.html)
	// in the Amazon RDS User Guide. Valid Values: IPV4 | DUAL
	NetworkType *string

	// The option group to associate the DB instance with. Permanent options, such as
	// the TDE option for Oracle Advanced Security TDE, can't be removed from an option
	// group. Also, that option group can't be removed from a DB instance after it is
	// associated with a DB instance. This setting doesn't apply to Amazon Aurora or
	// RDS Custom DB instances.
	OptionGroupName *string

	// The Amazon Web Services KMS key identifier for encryption of Performance
	// Insights data. The Amazon Web Services KMS key identifier is the key ARN, key
	// ID, alias ARN, or alias name for the KMS key. If you don't specify a value for
	// PerformanceInsightsKMSKeyId , then Amazon RDS uses your default KMS key. There
	// is a default KMS key for your Amazon Web Services account. Your Amazon Web
	// Services account has a different default KMS key for each Amazon Web Services
	// Region. This setting doesn't apply to RDS Custom DB instances.
	PerformanceInsightsKMSKeyId *string

	// The number of days to retain Performance Insights data. This setting doesn't
	// apply to RDS Custom DB instances. Valid Values:
	//   - 7
	//   - month * 31, where month is a number of months from 1-23. Examples: 93 (3
	//   months * 31), 341 (11 months * 31), 589 (19 months * 31)
	//   - 731
	// Default: 7 days If you specify a retention period that isn't valid, such as 94 ,
	// Amazon RDS returns an error.
	PerformanceInsightsRetentionPeriod *int32

	// The port number on which the database accepts connections. This setting doesn't
	// apply to Aurora DB instances. The port number is managed by the cluster. Valid
	// Values: 1150-65535 Default:
	//   - RDS for MariaDB - 3306
	//   - RDS for Microsoft SQL Server - 1433
	//   - RDS for MySQL - 3306
	//   - RDS for Oracle - 1521
	//   - RDS for PostgreSQL - 5432
	// Constraints:
	//   - For RDS for Microsoft SQL Server, the value can't be 1234 , 1434 , 3260 ,
	//   3343 , 3389 , 47001 , or 49152-49156 .
	Port *int32

	// The daily time range during which automated backups are created if automated
	// backups are enabled, using the BackupRetentionPeriod parameter. The default is
	// a 30-minute window selected at random from an 8-hour block of time for each
	// Amazon Web Services Region. For more information, see Backup window (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_WorkingWithAutomatedBackups.html#USER_WorkingWithAutomatedBackups.BackupWindow)
	// in the Amazon RDS User Guide. This setting doesn't apply to Amazon Aurora DB
	// instances. The daily time range for creating automated backups is managed by the
	// DB cluster. Constraints:
	//   - Must be in the format hh24:mi-hh24:mi .
	//   - Must be in Universal Coordinated Time (UTC).
	//   - Must not conflict with the preferred maintenance window.
	//   - Must be at least 30 minutes.
	PreferredBackupWindow *string

	// The time range each week during which system maintenance can occur. For more
	// information, see Amazon RDS Maintenance Window (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_UpgradeDBInstance.Maintenance.html#Concepts.DBMaintenance)
	// in the Amazon RDS User Guide. The default is a 30-minute window selected at
	// random from an 8-hour block of time for each Amazon Web Services Region,
	// occurring on a random day of the week. Constraints:
	//   - Must be in the format ddd:hh24:mi-ddd:hh24:mi .
	//   - The day values must be mon | tue | wed | thu | fri | sat | sun .
	//   - Must be in Universal Coordinated Time (UTC).
	//   - Must not conflict with the preferred backup window.
	//   - Must be at least 30 minutes.
	PreferredMaintenanceWindow *string

	// The number of CPU cores and the number of threads per core for the DB instance
	// class of the DB instance. This setting doesn't apply to Amazon Aurora or RDS
	// Custom DB instances.
	ProcessorFeatures []types.ProcessorFeature

	// The order of priority in which an Aurora Replica is promoted to the primary
	// instance after a failure of the existing primary instance. For more information,
	// see Fault Tolerance for an Aurora DB Cluster (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Concepts.AuroraHighAvailability.html#Aurora.Managing.FaultTolerance)
	// in the Amazon Aurora User Guide. This setting doesn't apply to RDS Custom DB
	// instances. Default: 1 Valid Values: 0 - 15
	PromotionTier *int32

	// Specifies whether the DB instance is publicly accessible. When the DB instance
	// is publicly accessible, its Domain Name System (DNS) endpoint resolves to the
	// private IP address from within the DB instance's virtual private cloud (VPC). It
	// resolves to the public IP address from outside of the DB instance's VPC. Access
	// to the DB instance is ultimately controlled by the security group it uses. That
	// public access is not permitted if the security group assigned to the DB instance
	// doesn't permit it. When the DB instance isn't publicly accessible, it is an
	// internal DB instance with a DNS name that resolves to a private IP address.
	// Default: The default behavior varies depending on whether DBSubnetGroupName is
	// specified. If DBSubnetGroupName isn't specified, and PubliclyAccessible isn't
	// specified, the following applies:
	//   - If the default VPC in the target Region doesn’t have an internet gateway
	//   attached to it, the DB instance is private.
	//   - If the default VPC in the target Region has an internet gateway attached to
	//   it, the DB instance is public.
	// If DBSubnetGroupName is specified, and PubliclyAccessible isn't specified, the
	// following applies:
	//   - If the subnets are part of a VPC that doesn’t have an internet gateway
	//   attached to it, the DB instance is private.
	//   - If the subnets are part of a VPC that has an internet gateway attached to
	//   it, the DB instance is public.
	PubliclyAccessible *bool

	// Specifes whether the DB instance is encrypted. By default, it isn't encrypted.
	// For RDS Custom DB instances, either enable this setting or leave it unset.
	// Otherwise, Amazon RDS reports an error. This setting doesn't apply to Amazon
	// Aurora DB instances. The encryption for DB instances is managed by the DB
	// cluster.
	StorageEncrypted *bool

	// The storage throughput value for the DB instance. This setting applies only to
	// the gp3 storage type. This setting doesn't apply to Amazon Aurora or RDS Custom
	// DB instances.
	StorageThroughput *int32

	// The storage type to associate with the DB instance. If you specify io1 or gp3 ,
	// you must also include a value for the Iops parameter. This setting doesn't
	// apply to Amazon Aurora DB instances. Storage is managed by the DB cluster. Valid
	// Values: gp2 | gp3 | io1 | standard Default: io1 , if the Iops parameter is
	// specified. Otherwise, gp2 .
	StorageType *string

	// Tags to assign to the DB instance.
	Tags []types.Tag

	// The ARN from the key store with which to associate the instance for TDE
	// encryption. This setting doesn't apply to Amazon Aurora or RDS Custom DB
	// instances.
	TdeCredentialArn *string

	// The password for the given ARN from the key store in order to access the
	// device. This setting doesn't apply to RDS Custom DB instances.
	TdeCredentialPassword *string

	// The time zone of the DB instance. The time zone parameter is currently
	// supported only by Microsoft SQL Server (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_SQLServer.html#SQLServer.Concepts.General.TimeZone)
	// .
	Timezone *string

	// A list of Amazon EC2 VPC security groups to associate with this DB instance.
	// This setting doesn't apply to Amazon Aurora DB instances. The associated list of
	// EC2 VPC security groups is managed by the DB cluster. Default: The default EC2
	// VPC security group for the DB subnet group's VPC.
	VpcSecurityGroupIds []string

	noSmithyDocumentSerde
}

type CreateDBInstanceOutput struct {

	// Contains the details of an Amazon RDS DB instance. This data type is used as a
	// response element in the operations CreateDBInstance ,
	// CreateDBInstanceReadReplica , DeleteDBInstance , DescribeDBInstances ,
	// ModifyDBInstance , PromoteReadReplica , RebootDBInstance ,
	// RestoreDBInstanceFromDBSnapshot , RestoreDBInstanceFromS3 ,
	// RestoreDBInstanceToPointInTime , StartDBInstance , and StopDBInstance .
	DBInstance *types.DBInstance

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDBInstanceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateDBInstance{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateDBInstance{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCreateDBInstanceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDBInstance(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCreateDBInstance(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "CreateDBInstance",
	}
}
