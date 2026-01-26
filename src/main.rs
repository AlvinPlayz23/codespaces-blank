use clap::Parser;
use log::{info, error};
use std::fs;
use std::path::PathBuf;

#[derive(Parser, Debug)]
#[command(author, version, about, long_about = None)]
struct Args {
    /// Path to the file to read
    #[arg(short, long)]
    file: PathBuf,
}

fn main() {
    // Initialize logger
    env_logger::init_from_env(env_logger::Env::default().default_filter_or("info"));

    let args = Args::parse();

    info!("Starting rusty_browser runner");
    info!("Reading file: {:?}", args.file);

    match fs::read_to_string(&args.file) {
        Ok(content) => {
            info!("Successfully read file. Size: {} bytes", content.len());
            // Diagnostic output
            println!("--- File Content Snippet ---");
            println!("{}", content.chars().take(100).collect::<String>());
            println!("----------------------------");
        }
        Err(e) => {
            error!("Failed to read file {:?}: {}", args.file, e);
            std::process::exit(1);
        }
    }
}
