# MLDatasetCreator\nA Golang tool that provides efficient creation of Machine Learning Datasets from public Social Network APIs. Its main function is to manage concurrent operations and save data, offering a user-friendly and intuitive API.\n\n## Overview\nSocial APIs are frequently used for creating datasets for training Machine Learning models. For instance, models like [Tweet2Vec](https://arxiv.org/abs/1607.07514) aim to extract features or create embeddings from such data.\n\nMany models, particularly NLP-oriented ones, can benefit from a large repository of structured text that may or may not carry labeling.\n\nOften, creating such datasets takes time away from feature engineering and model formulation work. This tool ai