require 'net/http'

class QiitaPostsController < ApplicationController

  before_action :find_qiita_post, only: %i(edit update destroy)

  DEPLOY_HOOKS_URI = URI.parse(ENV['PORTFOLIO_FRONT_DEPLOY_HOOKS_URL'])

  def index
    @qiita_posts = QiitaPost.order(created_at: :desc)
  end

  def new
    @qiita_post = QiitaPost.new
  end

  def create
    QiitaPost.create!(allowed_params)
    if Rails.env.production?
      Net::HTTP.get_response(DEPLOY_HOOKS_URI)
    end

    redirect_to qiita_posts_path
  rescue => e
    redirect_back fallback_location: qiita_posts_path
  end

  def edit; end

  def update
    @qiita_post.update!(allowed_params)
    if Rails.env.production?
      Net::HTTP.get_response(DEPLOY_HOOKS_URI)
    end

    redirect_to qiita_posts_path
  rescue => e
    redirect_back fallback_location: qiita_posts_path
  end

  def destroy
    @qiita_post.destroy!
    if Rails.env.production?
      Net::HTTP.get_response(DEPLOY_HOOKS_URI)
    end

    redirect_to qiita_posts_path
  rescue => e
    redirect_back fallback_location: qiita_posts_path
  end


  private

  def find_qiita_post
    @qiita_post = QiitaPost.find(params[:id])
  end

  def allowed_params
    params.require(:qiita_post).permit(
      :title,
      :url,
      :thumbnail_url,
      :category
    )
  end
end
