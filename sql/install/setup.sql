CREATE TABLE content_blocks (
	id INT PRIMARY KEY AUTO_INCREMENT,
	ext_id UUID NOT NULL DEFAULT UUID() UNIQUE,
	title VARCHAR(500),
	raw_content TEXT NOT NULL,
	raw_content_type VARCHAR(50) NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ,
  	updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP 
);

CREATE TABLE collections (
	id INT PRIMARY KEY AUTO_INCREMENT,
	ext_id UUID NOT NULL DEFAULT UUID() UNIQUE,
	name VARCHAR(500),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ,
  	updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP 
);


CREATE TABLE collection_content_blocks (
	id INT PRIMARY KEY AUTO_INCREMENT,
	ext_id UUID NOT NULL DEFAULT UUID() UNIQUE,
	collection_id INT NOT NULL,
	content_block_id INT NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT fk_collection_id FOREIGN KEY (collection_id) REFERENCES collections(id),
	CONSTRAINT fk_content_block_id FOREIGN KEY (content_block_id) REFERENCES content_blocks(id)
);

