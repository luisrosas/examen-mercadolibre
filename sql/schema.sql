CREATE TABLE `dnas` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `dna` varchar(255) NOT NULL,
  `is_mutant` tinyint(4) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1;
